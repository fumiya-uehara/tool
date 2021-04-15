package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

func NewToolCommand() *cobra.Command {
	cmds := &cobra.Command{
		Use: "tool",
	}
	cobra.OnInitialize(initConfig)

	cmds.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tool.yaml)")

	cmds.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cmds.AddCommand(NewCmdFfip())

	return cmds
}

func Execute() {
	cmd := NewToolCommand()
	cmd.SetOut(os.Stdout)
	if err := cmd.Execute(); err != nil {
		cmd.SetOut(os.Stderr)
		cmd.Println(err)
		os.Exit(1)
	}
}

func init() {
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".tool")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
