/*
Copyright © 2025 kyoruni <40832190+kyoruni@users.noreply.github.com>
*/
package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

//go:embed dict.json
var defaultDict []byte
var dictFile string

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
		fmt.Println("check called")

		var dictData []byte
		var err error

		if dictFile != "" {
			dictData, err = os.ReadFile(dictFile)
			if err != nil {
				fmt.Println("辞書ファイルの読み込みに失敗しました:", err)
				return
			}
			fmt.Println("辞書ファイルを読み込みました")
		} else {
			dictData = defaultDict
			fmt.Println("デフォルトの辞書ファイルを使用します")
		}

		fmt.Println("辞書ファイルの中身:")
		fmt.Println(string(dictData))
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	checkCmd.Flags().StringVarP(&dictFile, "dict", "d", "", "辞書ファイル(JSON)のパス")
}
