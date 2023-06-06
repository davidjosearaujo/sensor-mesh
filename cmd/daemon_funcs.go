/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"sensormesh/cmd/utils"
	"strings"
	"time"

	"berty.tech/go-orbit-db/iface"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var (
	MQTTmessageQueue []map[string]string
)

func publish() {
	defer cancel()
	lastTriggerTime := time.Now().Unix()
	currentTime := lastTriggerTime
	for {
		currentTime = time.Now().Unix()

		// whisper every 10 seconds
		if currentTime-lastTriggerTime >= 30 {
			lastTriggerTime = currentTime
			logger.Info().
				Str("type", "whisper").
				Str("name", utils.ViperConfs.GetString("name")).
				Send()

			// Posting new value to the log store
			_, err := logStore.Add(ctx, logbuf.Bytes())
			if err != nil {
				panic(fmt.Errorf("failed to put whisper in log store: %s", err))
			}

			// Reset reading buffer
			logbuf.Reset()
		}

		// Publishing MQTT messages to log store
		for len(MQTTmessageQueue) >= 1 {

			firstMessage := MQTTmessageQueue[0]

			// Convert map to JSON []byte string
			jsonMessage, err := json.Marshal(firstMessage)
			if err != nil {
				panic(fmt.Errorf("error converting MQTT message to json: %s", err))
			}

			// Posting new value to the log store
			_, err = logStore.Add(ctx, jsonMessage)
			if err != nil {
				panic(fmt.Errorf("failed to put MQTT message in log store: %s", err))
			}

			MQTTmessageQueue = MQTTmessageQueue[1:]
		}
	}
}

func subscribe() {
	defer cancel()
	var lastValue []byte
	file, err := os.OpenFile(utils.ViperConfs.GetString("logfile"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(fmt.Errorf("failed to open log file "+utils.ViperConfs.GetString("logfile")+": %s", err))
	}
	fileWriter := log.New(file, "", 0)
	defer file.Close()
	fmt.Println("[+] Writing log to " + utils.ViperConfs.GetString("logfile"))
	for {
		//Reading the last value inserted in the log store
		op, err := logStore.List(ctx, &iface.StreamOptions{})
		if err != nil {
			panic(fmt.Errorf("failed to get list from log store: %s", err))
		}

		if len(op) > 0 && !reflect.DeepEqual(op[0].GetValue(), lastValue) {
			//Write to log file if the new value is different from the last.
			//Since we are using timestamps, all correct messages will be
			//different, so this method becomes reliable in avoiding
			//incorrect or duplicate messages
			lastValue = op[0].GetValue()
			fileWriter.Println(strings.TrimRight(string(op[0].GetValue()), "\n"))
		}
	}
}

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	clientOptions := client.OptionsReader()

	newMessage := map[string]string{
		"sender":  clientOptions.ClientID(),
		"topic":   message.Topic(),
		"message": string(message.Payload()),
	}

	if len(MQTTmessageQueue) == 2 {
		fmt.Println("hi", len(MQTTmessageQueue))
	}

	MQTTmessageQueue = append(MQTTmessageQueue, newMessage)
}

// TODO - Periodically check if MQTT channels
// and topics have been changed, anf it so
// update connecting clients accordingly
