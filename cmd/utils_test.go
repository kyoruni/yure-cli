package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

// デフォルトの辞書から読み込み
func TestLoadDictFromEmbed(t *testing.T) {
	path := "test_dict.json"

	_, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("テスト用辞書ファイルの読み込みに失敗: %+v", err)
	}

	terms, err := loadDict(path, nil)
	if err != nil {
		t.Fatalf("テスト用辞書ファイルからの読み込みに失敗: %v", err)
	}

	if len(terms) != 1 {
		t.Errorf("読み込み件数が期待と違います: %+v", terms)
	}

	if terms[0].Wrong != "あっぷる" || terms[0].Correct != "アップル" {
		t.Errorf("読み込み結果が期待と違います: %+v", terms)
	}
}

// 辞書ファイルから読み込み
func TestLoadDictFromFile(t *testing.T) {
	content := `[{"correct":"テスト","wrong":"てすと"}]`
	tmpFile := "test_load_dict.json"
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatal("テスト用ファイルの作成に失敗:", err)
	}
	defer os.Remove(tmpFile) // 終わったら削除

	terms, err := loadDict(tmpFile, nil)
	if err != nil {
		t.Fatalf("ファイルからの読み込みに失敗: %v", err)
	}

	if len(terms) != 1 {
		t.Fatalf("読み込み件数が期待と違います: %+v", terms)
	}

	if terms[0].Wrong != "てすと" || terms[0].Correct != "テスト" {
		t.Errorf("読み込み内容が期待と違います: %+v", terms)
	}
}

// 入力ファイルから読み込み
func TestLoadInputFile(t *testing.T) {
	content := "これはテストファイルです\n2行目です\n"
	tmpFile := "test_input.txt"

	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatal("テスト用ファイルの作成に失敗", err)
	}
	defer os.Remove(tmpFile) // 終わったら削除

	data, err := loadInputFile(tmpFile)
	if err != nil {
		t.Fatalf("ファイルからの読み込みに失敗: %v", err)
	}

	if string(data) != content {
		t.Errorf("読み込み結果が期待と違います\n期待: %q\n実際: %q", content, string(data))
	}
}

// NGワードを探す
func TestFindWrongTerms(t *testing.T) {
	lines := []string{
		"これは正しい文です",
		"これはNginxを使った行です",
		"これはピカチューが出てくる行です",
		"これは無関係な行です",
	}
	terms := []Term{
		{Correct: "nginx", Wrong: "Nginx"},
		{Correct: "ピカチュウ", Wrong: "ピカチュー"},
	}

	// 標準出力をキャプチャする
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	findWrongTerms(lines, terms, "testfile.txt")

	w.Close()
	os.Stdout = stdout
	_, _ = buf.ReadFrom(r)

	output := buf.String()

	// 検証する
	expected := []string{
		"testfile.txt:2: これはNginxを使った行です",
		"testfile.txt:3: これはピカチューが出てくる行です",
	}

	for _, e := range expected {
		if !strings.Contains(output, e) {
			t.Errorf("出力に期待した行が見つかりませんでした: %q\n出力全体:\n%s", e, output)
		}
	}
}
