/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>
*/
package utils

import (
	"fmt"
)

func AddChannel(broker string, topics []string) {

	newChannel := map[string]interface{}{
		"broker": broker,
		"topics": topics,
	}

	// Check if list is empty
	channelList := []interface{}{}
	if ViperConfs.IsSet("channels") {
		// Initialize "channels" key with an empty slice
		channelList = ViperConfs.Get("channels").([]interface{})
	}

	channelList = append(channelList, newChannel)
	ViperConfs.Set("channels", channelList)
	err := ViperConfs.WriteConfig()
	if err != nil {
		panic(fmt.Errorf("error adding new channel: %v", err))
	}
}

func DisconnectChannelTopic(broker string, topics []string) {
	if ViperConfs.Get("channels") == nil {
		panic(fmt.Errorf("there are no channels to remove"))
	}

	channelList := ViperConfs.Get("channels").([]interface{})
	updatedChannels := make([]interface{}, 0, len(channelList))

	for _, channel := range channelList {
		channelMap := channel.(map[string]interface{})
		if channelMap["broker"].(string) != broker {
			updatedChannels = append(updatedChannels, channel)
		} else if len(topics) > 0 {
			topicList := channelMap["topics"].([]interface {})

			removingTopics := make(map[string]bool)
			updatedTopics := []string{}

			for _, removingTopic := range topics {
				removingTopics[removingTopic] = true
			}

			for _, existingTopic := range topicList {
				if _, found := removingTopics[existingTopic.(string)]; !found {
					updatedTopics = append(updatedTopics, existingTopic.(string))
				}
			}

			modifiedChannel := map[string]interface{}{
				"broker": channelMap["broker"].(string),
				"topics": updatedTopics,
			}
			updatedChannels = append(updatedChannels, modifiedChannel)
		}

	}

	ViperConfs.Set("channels", updatedChannels)
	err := ViperConfs.WriteConfig()
	if err != nil {
		if len(topics) != 0 {
			panic(fmt.Errorf("error removing topic: %v", err))
		}
		panic(fmt.Errorf("error removing channel: %v", err))
	}
}