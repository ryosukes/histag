package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "histag",
	Short: "This tool is pretty cool.",
	Long:  "This tool is a great convenience.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(catCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of histag",
	Long:  `All software has versions. This is histag's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("histag v1.0")
	},
}

var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "Print the version number of histag",
	Long:  `All software has versions. This is histag's`,
	Run: func(cmd *cobra.Command, args []string) {
		histfile := os.Getenv("HISTFILE")
		history, err := ioutil.ReadFile(histfile)
		if err != nil {
			// エラー処理
		}
		fmt.Println(string(history))
	},
}
