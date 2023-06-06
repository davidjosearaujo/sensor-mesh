/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>
*/
package sensor

import (
	"fmt"
	"sensormesh/cmd/utils"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add a new sensor",
	Example: "sensormesh sensor add --name=\"humidity\" --baud=9600 --parity=\"None\" --port=\"/dev/ttyUSB0\" --size=8 --stop=\"Stop1\" --interval=60",
	PreRun: func(cmd *cobra.Command, args []string) {
		utils.LoadConfigurationFromFile()
	},
	Run: func(cmd *cobra.Command, args []string) {
		utils.AddSensor(_name, _port, _baud, _size, _parity, _stop, _interval)
		fmt.Println("[!] Sensor add successfully!")
	},
}

func init() {
	SensorCmd.AddCommand(addCmd)

	addCmd.Flags().StringVar(&_name, "name", "Sensor 1", "Name of the sensor to be connected. Defaults to \"Sensor 1\"")
	_ = addCmd.MarkFlagRequired("name")

	addCmd.Flags().StringVar(&_port, "port", "/dev/ttyUSB0", "Connection port to be used for the sensor. Defaults to \"/dev/ttyUSB0\"")
	_ = addCmd.MarkFlagRequired("port")

	addCmd.Flags().IntVar(&_baud, "baud", 9600, "Baud rate of the sensor. Defaults to 9600")
	_ = addCmd.MarkFlagRequired("baud")

	addCmd.Flags().IntVar(&_size, "size", 8, "Character size in transmission, commonly either 7 or 8. Defaults to 8")

	addCmd.Flags().StringVar(&_parity, "parity", "None", "Error detecting parity bit mode, commonly either Even, Odd, None, Mark(always 1) or Space(always 0). Defaults to \"None\"")

	addCmd.Flags().StringVar(&_stop, "stop", "Stop1", "End or character transmission signal bits, commonly wither Stop1(1), Stop2(2) or Stop1Half(15). Defaults to \"Stop1\"")

	addCmd.Flags().IntVar(&_interval, "interval", 60, "Interval in seconds between sensor queries. Defaults to 60")

	addCmd.MarkFlagsRequiredTogether("name", "port", "baud", "size")
	addCmd.MarkFlagsRequiredTogether("name", "port", "baud", "parity")
	addCmd.MarkFlagsRequiredTogether("name", "port", "baud", "stop")
}
