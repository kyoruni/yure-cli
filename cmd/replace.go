/*
Copyright © 2025 kyoruni <40832190+kyoruni@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/kyoruni/yure-cli/embeddata"
	"github.com/spf13/cobra"
)

var replaceDictFile string
var replaceInputFile string

// replaceCmd represents the replace command
var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "表記ゆれを修正して上書き保存します",
	Run: func(cmd *cobra.Command, args []string) {
		terms, err := loadDict(replaceDictFile, embeddata.GetDefaultDict())
		if err != nil {
			fmt.Println("辞書ファイルの読み込みに失敗しました:", err)
			return
		}

		if replaceInputFile == "" {
			fmt.Println("入力ファイルが指定されていません")
			return
		}

		content, err := loadInputFile(replaceInputFile)
		if err != nil {
			fmt.Println(err)
			return
		}

		replaced := replaceTerms(content, terms)

		output := strings.Join(replaced, "\n")
		err = os.WriteFile(replaceInputFile, []byte(output), 0644)
		if err != nil {
			fmt.Println("ファイルの上書き保存に失敗しました:", err)
			return
		}

		fmt.Printf("置換した内容を %s に保存しました\n", replaceInputFile)
	},
}

func replaceTerms(lines []string, terms []Term) []string {
	var replaced []string
	for _, line := range lines {
		newLine := line
		for _, term := range terms {
			newLine = strings.ReplaceAll(newLine, term.Wrong, term.Correct)
		}
		replaced = append(replaced, newLine)
	}
	return replaced
}

func init() {
	rootCmd.AddCommand(replaceCmd)

	replaceCmd.Flags().StringVarP(&replaceDictFile, "dict", "d", "", "辞書ファイル(JSON)のパス")
	replaceCmd.Flags().StringVarP(&replaceInputFile, "input", "i", "", "チェック対象のテキストファイル")
}
