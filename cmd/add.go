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

var comment string

func add(cmd *cobra.Command, args []string) {
	histfile := os.Getenv("HISTFILE")

	history, err := ioutil.ReadFile(histfile)
	if err != nil {
		// エラー処理
	}

	line, _ := filter(string(history))

	fmt.Println("Please enter a comment for " + "'" + line + "'")
	fmt.Scan(&comment)

	text := fmt.Sprintf("%s #%s", line, comment)

	file, _ := os.OpenFile(histfile, os.O_APPEND|os.O_WRONLY, 0600)
	defer file.Close()

	fmt.Fprintln(file, text)
	fmt.Println("add command to $HISTFILE!: " + text)
}
