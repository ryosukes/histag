package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Print the version number of histag",
	Long:  `All software has versions. This is histag's`,
	Run:   add,
}

func add(cmd *cobra.Command, args []string) {
	histfile := os.Getenv("HISTFILE")
	f, _ := os.OpenFile(histfile, os.O_APPEND|os.O_WRONLY, 0600)
	defer f.Close()

	history, err := ioutil.ReadFile(histfile)

	content := []byte("#add test")
	ioutil.WriteFile(histfile, content, os.ModePerm)
	if err != nil {
		// エラー処理
	}
	fmt.Println(string(history))
}
