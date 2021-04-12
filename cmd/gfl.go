package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	gflCmd = &cobra.Command{
		Use: "gfl",
		RunE: func(cmd *cobra.Command, args []string) error {
			//			files, err := ioutil.ReadDir("./test")
			//			if err != nil {
			//				panic(err)
			//			}
			//
			//			for _, f := range files {
			//				fmt.Println(f.Name())
			//			}

			err := filepath.Walk("./test", func(path string, info os.FileInfo, err error) error {
				if err != nil {
					panic(err)
				}
				if !info.IsDir() {
					fmt.Printf("path: %#v\n", path)
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
}
