/*
Copyright © 2025 kyoruni <40832190+kyoruni@users.noreply.github.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var replaceDictFile string
var replaceInputFile string

// replaceCmd represents the replace command
var replaceCmd = &cobra.Command{
	Use:   "replace",
	Short: "表記揺れを修正して上書き保存します",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := loadDict(replaceDictFile, defaultDict)
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

		fmt.Println("入力ファイルの内容:")
		for i, line := range content {
			fmt.Printf("%2d: %s\n", i+1, line)
		}
	},
}

func init() {
	rootCmd.AddCommand(replaceCmd)

	replaceCmd.Flags().StringVarP(&replaceDictFile, "dict", "d", "", "辞書ファイル(JSON)のパス")
	replaceCmd.Flags().StringVarP(&replaceInputFile, "input", "i", "", "チェック対象のテキストファイル")
}
