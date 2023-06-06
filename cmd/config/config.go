/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>
*/
package config

import (
	"fmt"
	"os"
	"sensormesh/cmd/utils"
	"strings"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Config allows you to see your configurations and edit them",
	PreRun: func(cmd *cobra.Command, args []string) {
		utils.LoadConfigurationFromFile()
	},
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 2:
			utils.ViperConfs.Set(args[0], args[1])
		case 1:
			fmt.Println(utils.ViperConfs.Get(args[0]))
		default:
			contents, err := os.ReadFile(utils.ConfigFilePath)
			if err != nil {
				fmt.Println("File reading error", err)
				return
			}
			fmt.Println(strings.TrimRight(string(contents), "\n"))
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		utils.ViperConfs.WriteConfig()
	},
}

func init() {
	ConfigCmd.DisableFlagsInUseLine = true
}
