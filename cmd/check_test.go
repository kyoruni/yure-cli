package cmd

import (
	"os"
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
