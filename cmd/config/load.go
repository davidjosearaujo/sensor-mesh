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
	"firemesh/cmd/shared"
	"github.com/spf13/cobra"
)

// loadCmd represents the load command
var loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Loads path's given file as default configuration",
	Run: func(cmd *cobra.Command, args []string) {
		shared.LoadConfigurationFromFile()
	},
}

func init() {
	ConfigCmd.AddCommand(loadCmd)

	loadCmd.Flags().StringVar(&shared.ConfigFilePath, "path", shared.ConfigFilePath, "FireMesh configuration file. Defaults to \"~/.firemesh/config.yaml\"")
	_ = loadCmd.MarkFlagRequired("path")
}
