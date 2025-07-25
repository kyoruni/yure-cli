/*
Copyright © 2025 kyoruni <40832190+kyoruni@users.noreply.github.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Term struct {
	Correct string `json:"correct"`
	Wrong   string `json:"wrong"`
}

func loadDict(dictFile string, embedded []byte) ([]Term, error) {
	var dictData []byte
	var err error

	if dictFile != "" {
		dictData, err = os.ReadFile(dictFile)
		if err != nil {
			return nil, fmt.Errorf("辞書ファイルの読み込みに失敗しました: %w", err)
		}
	} else {
		dictData = embedded
	}

	var terms []Term
	err = json.Unmarshal(dictData, &terms)
	if err != nil {
		return nil, fmt.Errorf("辞書ファイルの展開に失敗しました: %w", err)
	}

	return terms, nil
}

func loadInputFile(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("入力ファイルの読み込みに失敗しました: %w", err)
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func findWrongTerms(lines []string, terms []Term, fileName string) {
	for i, line := range lines {
		for _, term := range terms {
			if strings.Contains(line, term.Wrong) {
				fmt.Printf("%s:%d: %s\n", fileName, i+1, line)
				break
			}
		}
	}
}
