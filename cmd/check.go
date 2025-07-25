/*
Copyright © 2025 kyoruni <40832190+kyoruni@users.noreply.github.com>
*/
package cmd

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/kyoruni/yure-cli/embeddata"
	"github.com/spf13/cobra"
)

var dictFile string
var inputFile string

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "表記ゆれに該当する箇所を探します",
	Run: func(cmd *cobra.Command, args []string) {
		terms, err := loadDict(dictFile, embeddata.GetDefaultDict())
		if err != nil {
			fmt.Println("辞書ファイルの読み込みに失敗しました:", err)
			return
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

		findWrongTerms(content, terms, inputFile)
	},
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

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringVarP(&dictFile, "dict", "d", "", "辞書ファイル(JSON)のパス")
	checkCmd.Flags().StringVarP(&inputFile, "input", "i", "", "チェック対象のテキストファイル")
}
