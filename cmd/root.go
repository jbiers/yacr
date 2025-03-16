package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yacr",
	Short: "Yet Another Container Runtime",
	Long: `A runc-inspired container runtime. 
Built with the purpose of peeling away the abstractions to understand how a container works under the hood.`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(createCommand)
	rootCmd.AddCommand(deleteCommand)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
