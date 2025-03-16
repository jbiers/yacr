package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCommand = &cobra.Command{
	Use:   "delete",
	Short: "Delete an existing container",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Deleting container...")
	},
}
