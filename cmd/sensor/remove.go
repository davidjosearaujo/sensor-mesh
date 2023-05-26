/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>
*/
package sensor

import (
	"sensormesh/cmd/shared"

	"github.com/spf13/cobra"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:     "remove",
	Short:   "Remove a sensor",
	Example: "sensormesh sensor remove --name=\"humidity\"",
	Run: func(cmd *cobra.Command, args []string) {
		shared.RemoveSensor(_name)
	},
}

func init() {
	SensorCmd.AddCommand(removeCmd)

	removeCmd.Flags().StringVar(&_name, "name", "", "Name of the sensor to be removed")
	_ = removeCmd.MarkFlagRequired("name")
}
