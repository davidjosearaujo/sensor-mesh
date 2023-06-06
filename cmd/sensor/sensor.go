/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>
*/
package sensor

import (
	"github.com/spf13/cobra"
)

var (
	_name     string
	_port     string
	_baud     int
	_size     int
	_parity   string
	_stop     string
	_interval int
)

// sensorCmd represents the sensor command
var SensorCmd = &cobra.Command{
	Use:   "sensor",
	Short: "Sensor is a palette that contains sensor based commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	SensorCmd.DisableFlagsInUseLine = true
}
