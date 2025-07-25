package cmd

import (
	"testing"
)

func TestReplaceTerms(t *testing.T) {
	lines := []string{
		"これはNginxを使った行です",
		"ピカチューが現れた",
		"これは正しい文ですピカチュウnginx",
	}
	terms := []Term{
		{Correct: "nginx", Wrong: "Nginx"},
		{Correct: "ピカチュウ", Wrong: "ピカチュー"},
	}

	expected := []string{
		"これはnginxを使った行です",
		"ピカチュウが現れた",
		"これは正しい文ですピカチュウnginx",
	}

	result := replaceTerms(lines, terms)

	if len(result) != len(expected) {
		t.Fatalf("出力の行数が違います。期待: %d, 実際: %d", len(expected), len(result))
	}

	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("行 %d の置換結果が期待と異なります\n期待: %q\n実際: %q", i+1, expected[i], result[i])
		}
	}
}
