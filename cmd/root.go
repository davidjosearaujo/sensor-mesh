/*
Copyright © 2023 David Araújo <davidaraujo98.github.io>
*/
package cmd

import (
	"os"
	"sensormesh/cmd/channel"
	"sensormesh/cmd/config"
	"sensormesh/cmd/sensor"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sensormesh",
	Short: "",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.DisableFlagsInUseLine = true
	rootCmd.AddCommand(
		sensor.SensorCmd,
		config.ConfigCmd,
		channel.ChannelCmd,
	)
}
