/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>
*/
package channel

import (
	"fmt"
	"sensormesh/cmd/utils"

	"github.com/spf13/cobra"
)

// disconnectCmd represents the disconnect command
var disconnectCmd = &cobra.Command{
	Use:     "disconnect",
	Short:   "Disconnect from a MQTT channel topic. If no topic specified, disconnects entirely from broker.",
	Example: "sensormesh channel disconnect --brokerUrl=\"tcp://mqtt.example.com:1883\" --topic=\"topic/topic\" --topic=\"topic/topic2\" ...",
	PreRun: func(cmd *cobra.Command, args []string) {
		utils.LoadConfigurationFromFile()
	},
	Run: func(cmd *cobra.Command, args []string) {
		utils.DisconnectChannelTopic(_brokerUrl, _topics)
		if len(_topics) > 0 {
			fmt.Println("[!] Unsubscribed from channel successfully! If SensorMesh daemon running, it must be rebooted to take effect!")
			return
		}
		fmt.Println("[!] Disconnected from channel successfully! If SensorMesh daemon running, it must be rebooted to take effect!")
	},
}

func init() {
	ChannelCmd.AddCommand(disconnectCmd)

	disconnectCmd.Flags().StringVar(&_brokerUrl, "brokerUrl", "tcp://mqtt.example.com:1883", "MQTT broker url. Defaults to \"tcp://localhost:1883\"")
	_ = disconnectCmd.MarkFlagRequired("brokerUrl")

	disconnectCmd.Flags().StringArrayVar(&_topics, "topic", []string{}, "Topics to subscribe. Defaults to \"topic/sensormesh\"")
}
