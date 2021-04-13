package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	src string

	gflCmd = &cobra.Command{
		Use: "gfl",
		RunE: func(cmd *cobra.Command, args []string) error {
			s, err := cmd.Flags().GetString("source")
			if err != nil {
				panic(err)
			}
			err = filepath.Walk(s, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					panic(err)
				}
				if !info.IsDir() {
					fmt.Println(path)
				}
				return nil

			})
			if err != nil {
				panic(err)
			}
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(gflCmd)
	gflCmd.Flags().StringVarP(&src, "source", "s", "", "source directory")
	_ = gflCmd.MarkFlagRequired("source")
}
