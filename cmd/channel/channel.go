/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>
*/
package channel

import (
	"github.com/spf13/cobra"
)

var (
	_brokerUrl string
	_topics    []string
)

// channelCmd represents the channel command
var ChannelCmd = &cobra.Command{
	Use:   "channel",
	Short: "Channel is a palette that contains MQTT addion based commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	ChannelCmd.DisableFlagsInUseLine = true
}
