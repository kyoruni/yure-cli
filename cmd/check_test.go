package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

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
