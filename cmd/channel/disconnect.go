/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package channel

import (
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
	},
}

func init() {
	ChannelCmd.AddCommand(disconnectCmd)

	disconnectCmd.Flags().StringVar(&_brokerUrl, "brokerUrl", "tcp://mqtt.example.com:1883", "MQTT broker url. Defaults to \"tcp://localhost:1883\"")
	_ = disconnectCmd.MarkFlagRequired("brokerUrl")

	disconnectCmd.Flags().StringArrayVar(&_topics, "topic", []string{}, "Topics to subscribe. Defaults to \"topic/sensormesh\"")
}
