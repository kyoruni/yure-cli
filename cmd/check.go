/*
Copyright © 2025 kyoruni <40832190+kyoruni@users.noreply.github.com>
*/
package cmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Term struct {
	Correct string `json:"correct"`
	Wrong   string `json:"wrong"`
}

//go:embed dict.json
var defaultDict []byte
var dictFile string
var inputFile string

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		terms, err := loadDict(dictFile, defaultDict)
		if err != nil {
			fmt.Println("辞書ファイルの読み込みに失敗しました:", err)
			return
		}

		fmt.Println("辞書ファイルの中身:")
		for _, t := range terms {
			fmt.Printf(" NG %s => OK %s\n", t.Wrong, t.Correct)
		}

		if inputFile == "" {
			fmt.Println("入力ファイルが指定されていません")
			return
		}

		content, err := loadInputFile(inputFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("\n---入力ファイルの内容---")
		fmt.Println(string(content))

		lines := strings.Split(string(content), "\n")
		for i, line := range lines {
			for _, term := range terms {
				if strings.Contains(line, term.Wrong) {
					fmt.Printf("%s:%d: %s\n", inputFile, i+1, line)
					break
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringVarP(&dictFile, "dict", "d", "", "辞書ファイル(JSON)のパス")
	checkCmd.Flags().StringVarP(&inputFile, "input", "i", "", "チェック対象のテキストファイル")
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

func loadInputFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("入力ファイルの読み込みに失敗しました: %w", err)
	}
	return data, nil
}
