package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCommand = &cobra.Command{
	Use:   "create",
	Short: "Create a new container",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Creating container...")
	},
}
