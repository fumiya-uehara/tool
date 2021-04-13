package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type path string

var (
	wg sync.WaitGroup

	src string
	threads int

	gflCmd = &cobra.Command{
		Use: "gfl",
		RunE: func(cmd *cobra.Command, args []string) error {

			s, err := cmd.Flags().GetString("source")
			if err != nil {
				panic(err)
			}

			files, err := ioutil.ReadDir(s)
			if err != nil {
				panic(err)
			}
			var divided [][]fs.FileInfo
			chunkSize := len(files)/threads-1
			for i:=0; i<len(files); i+=chunkSize {
				end := i + chunkSize

				if end > len(files) {
					end = len(files)
				}

				divided = append(divided, files[i:end])
			}

			for _, infoList := range divided {
				wg.Add(1)
				go func(il []fs.FileInfo) {
					defer wg.Done()
					for _, info := range il {
						fetchPathForTypeIsFile(s, info)
					}
				}(infoList)
			}
			wg.Wait()
			return nil
		},
	}
)

func fetchPathForTypeIsFile(rootPath string, info fs.FileInfo) {
	path := strings.Join([]string{rootPath, info.Name()}, "/")
	err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}
		if !d.IsDir() {
			fmt.Println(path)
		}
		return nil

	})
	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(gflCmd)
	gflCmd.Flags().StringVarP(&src, "source", "s", "", "source directory")
	_ = gflCmd.MarkFlagRequired("source")

	gflCmd.Flags().IntVarP(&threads, "threads", "p", 1, "concurrency count")
}
