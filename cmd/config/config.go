/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package config

import (
	"fmt"
	"os"
	"sensormesh/cmd/shared"
	"strings"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Config allows you to see your configurations and edit them",
	PreRun: func(cmd *cobra.Command, args []string) {
		shared.LoadConfigurationFromFile()
	},
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 2:
			shared.ViperConfs.Set(args[0], args[1])
		case 1:
			fmt.Println(shared.ViperConfs.Get(args[0]))
		default:
			contents, err := os.ReadFile(shared.ConfigFilePath)
			if err != nil {
				fmt.Println("File reading error", err)
				return
			}
			fmt.Println(strings.TrimRight(string(contents), "\n"))
		}
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		shared.ViperConfs.WriteConfig()
	},
}

func init() {
	ConfigCmd.DisableFlagsInUseLine = true
}
