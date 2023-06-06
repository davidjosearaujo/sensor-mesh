/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>
*/
package cmd

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sensormesh/cmd/utils"

	orbitdb "berty.tech/go-orbit-db"
	"berty.tech/go-orbit-db/accesscontroller"
	"berty.tech/go-orbit-db/iface"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	client "github.com/ipfs/go-ipfs-http-client"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	MQTTClients  []MQTT.Client
	ctx          context.Context
	cancel       context.CancelFunc
	storeAddress string
	logbuf       bytes.Buffer
	logger       zerolog.Logger
	logStore     iface.EventLogStore
)

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run a OrbitDB sensor logger",
	Long: `'sensormesh daemon' runs a persistent sensormesh daemon that can
query specified sensor and log their responses to a OrbitDB
log file, that will be utils between nodes in a same IPFS
private network.

The daemon will start by first configuring the current
machine as a node in a private IPFS network, and then
initialize IPFS's daemon`,
	PreRun: func(cmd *cobra.Command, args []string) {
		// Check existence of config file
		_, err := utils.Exists(utils.ConfigFilePath)
		if err != nil {
			panic(fmt.Errorf("configuration file not set. Try running 'sensormesh init' first: %s", err))
		}

		// Load configurations from the configurations file, if non-existing, a new one will be created
		utils.LoadConfigurationFromFile()

		// Connecting to local IPFS node API
		shell, err := client.NewURLApiWithClient(utils.LocalIPFSApiAddress(), &http.Client{})
		if err != nil {
			panic(fmt.Errorf("failed to connect to local IPFS API. IPFS daemon must be running with '--enable-pubsub-experiment': %s", err))
		}
		fmt.Printf("[+] Connecting to %s's local IPFS API at %s\n", utils.ViperConfs.GetString("name"), utils.LocalIPFSApiAddress())

		ctx, cancel = context.WithCancel(context.Background())

		// Initiating a new OrbitDB instance
		db, err := orbitdb.NewOrbitDB(ctx, shell, &orbitdb.NewOrbitDBOptions{})
		if err != nil {
			panic(fmt.Errorf("failed to create new orbitdb. IPFS daemon must be running with '--enable-pubsub-experiment': %s", err))
		}

		// Give read and write permissions to all
		ac := &accesscontroller.CreateAccessControllerOptions{
			Access: map[string][]string{
				"write": {
					"*",
				},
				"read": {
					"*",
				},
			},
		}

		// Tries connecting to the given address (If name was give, error will trigger creation on new database)
		logStore, err = db.Log(ctx, storeAddress, &orbitdb.CreateDBOptions{
			AccessController: ac,
		})
		if err != nil {
			panic(fmt.Errorf("failed to get log store: %s", err))
		}
		fmt.Printf("[!] Connected to %s store with address: %s\n", logStore.DBName(), logStore.Address().String())
		utils.ViperConfs.Set("orbitdb.storeaddress", logStore.Address().String())

		// Set up MQTT clients
		if utils.ViperConfs.IsSet("channels") {
			// Initialize "channels" key with an empty slice
			for _, channel := range utils.ViperConfs.Get("channels").([]interface{}) {
				channelMap := channel.(map[string]interface{})
				opts := MQTT.NewClientOptions()
				opts.AddBroker(channelMap["broker"].(string))
				opts.SetClientID(utils.ViperConfs.GetString("name"))

				// Create and start a new MQTT client
				client := MQTT.NewClient(opts)
				if token := client.Connect(); token.Wait() && token.Error() != nil {
					panic(fmt.Errorf("error connecting to MQTT: %v", token.Error()))
				}
				fmt.Printf("[!] Connected to %s MQTT channel with client ID of : %s\n", channelMap["broker"].(string), utils.ViperConfs.GetString("name"))

				// Subscribe to the topic
				for _, topic := range channelMap["topics"].([]interface{}) {
					// Subscribe to the topic
					if token := client.Subscribe(topic.(string), 0, onMessageReceived); token.Wait() && token.Error() != nil {
						panic(fmt.Errorf("error subscribing to topic in MQTT: %v", token.Error()))
					}
					fmt.Printf("[+] Subscribed to %s MQTT topic\n", topic.(string))
				}
			}
		}

		// Initialize zerolog logger
		logger = zerolog.New(&logbuf).With().Timestamp().Logger()

		err = utils.ViperConfs.WriteConfig()
		if err != nil {
			panic(fmt.Errorf("error updating config file: %v", err))
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Initiate reading and writing to the database as a multi-threaded processes
		go publish()
		go subscribe()

		fmt.Println("[+] Press Ctrl+c to stop daemon")

		// Capture SIGINT
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)

		// Wait for either WaitGroup or interrupt signal
		<-sigint

		for _, client := range MQTTClients {
			client.Disconnect(0)
		}

		fmt.Println("[!!!] Interrupt signal received, terminating...")
	},
}

func init() {
	rootCmd.AddCommand(daemonCmd)
	daemonCmd.Flags().StringVar(&storeAddress, "storeaddress", "event", "Address of the log store. Defaults to create a new log store with name 'event'")
	_ = daemonCmd.MarkFlagRequired("name")
}
