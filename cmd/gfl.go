package cmd

import (
	"github.com/spf13/cobra"
)

var (
	gflCmd = &cobra.Command{
		Use: "gfl",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(gflCmd)
}
