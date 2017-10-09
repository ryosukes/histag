package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCmd is root for histag
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
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of histag",
	Long:  `All software has versions. This is histag's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("histag v1.0")
	},
}
