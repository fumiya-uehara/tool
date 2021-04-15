package cmd

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

func NewCmdFfip() *cobra.Command {
	var (
		wg      sync.WaitGroup
		threads int
	)

	cmd := &cobra.Command{
		Use: "ffip",
		Args: func(cmd *cobra.Command, args []string) error {
			// FIXME 一旦引数は一つに限定する
			if len(args) != 1 {
				return errors.New("only one required a target path")
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// TODO 現状は引数一つのみを許容しているのでargs[0]の実装になっている
			files, err := ioutil.ReadDir(args[0])
			if err != nil {
				panic(err)
			}

			var divided [][]fs.FileInfo
			chunkSize := len(files)/threads - 1
			for i := 0; i < len(files); i += chunkSize {
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
						fetchPathForTypeIsFile(cmd, args[0], info)
					}
				}(infoList)
			}
			wg.Wait()
			return nil
		},
	}

	cmd.Flags().IntVarP(&threads, "threads", "p", 1, "concurrency count")

	return cmd
}

func fetchPathForTypeIsFile(c *cobra.Command, rootPath string, info fs.FileInfo) {
	path := strings.Join([]string{rootPath, info.Name()}, "/")
	err := filepath.WalkDir(path, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			panic(err)
		}
		if !d.IsDir() {
			c.Println(path)
		}
		return nil

	})
	if err != nil {
		panic(err)
	}
}

func init() {
}
