/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package channel

import (
	"sensormesh/cmd/utils"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var connectCmd = &cobra.Command{
	Use:     "connect",
	Short:   "Connect to a new MQTT channel topic",
	Example: "sensormesh channel connect --brokerUrl=\"tcp://mqtt.example.com:1883\" --topic=\"topic/topic\" --topic=\"topic/topic2\" ...",
	PreRun: func(cmd *cobra.Command, args []string) {
		utils.LoadConfigurationFromFile()
	},
	Run: func(cmd *cobra.Command, args []string) {
		utils.AddChannel(_brokerUrl, _topics)
	},
}

func init() {
	ChannelCmd.AddCommand(connectCmd)

	connectCmd.Flags().StringVar(&_brokerUrl, "brokerUrl", "tcp://mqtt.example.com:1883", "MQTT broker url. Defaults to \"tcp://localhost:1883\"")
	_ = connectCmd.MarkFlagRequired("brokerUrl")

	connectCmd.Flags().StringArrayVar(&_topics, "topic", []string{"topic/sensormesh"}, "Topics to subscribe. Defaults to \"topic/sensormesh\"")
	_ = connectCmd.MarkFlagRequired("topic")
}
