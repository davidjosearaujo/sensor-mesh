/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>
*/
package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sensormesh/cmd/shared"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/spf13/cobra"
)

var (
	IpfsApi          *shell.Shell
	IpfsPath         string
	swarmKey         string
	nodename         string
	logfile          string
	swarmKeyFilePath = filepath.Join(shared.GetUserHomeDir(), ".ipfs", "swarm.key")
)

func createSwarmKeyFile() {
	// Clear old swarm key file
	os.Remove(swarmKeyFilePath)

	// If swarm key not provided, creates new one
	if swarmKey == "" {
		key := make([]byte, 32)
		_, err := rand.Read(key)
		if err != nil {
			fmt.Println("While trying to read random source for swarm key:", err)
			os.Exit(1)
		}
		swarmKey = hex.EncodeToString(key)
		fmt.Println("New swarm key: ", swarmKey)
	}

	var (
		file *os.File
		err  error
	)
	exists, _ := shared.Exists(swarmKeyFilePath)
	if !exists {
		err = os.MkdirAll(filepath.Dir(swarmKeyFilePath), 0700)
		if err != nil {
			fmt.Printf("Error creating directories: %v\n", err)
			os.Exit(1)
		}
		file, err = os.Create(swarmKeyFilePath)
		if err != nil {
			fmt.Printf("Error generating swarm file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		fmt.Fprintln(file, "/key/swarm/psk/1.0.0/")
		fmt.Fprintln(file, "/base16/")
		fmt.Fprintln(file, swarmKey)
	}
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize local FireMesh configuration",
	Run: func(cmd *cobra.Command, args []string) {
		// Checking if IPFS configs exist
		_, err := shared.Exists(filepath.Join(shared.GetUserHomeDir(), ".ipfs", "config"))
		if err != nil {
			panic(fmt.Errorf("configuration file not set. Try running 'ipfs init' first: %s", err))
		}

		// Creates new swarm key file
		createSwarmKeyFile()

		// Changing to DHT type routing
		ipfsCmd := exec.Command("ipfs", "config", "Routing.Type", "dht")
		err = ipfsCmd.Run()
		if err != nil {
			panic(fmt.Errorf("error while starting IPFS config files: %v", err))
		}

		// Remove all bootstrap addresses
		ipfsCmd = exec.Command("ipfs", "bootstrap", "rm", "--all")
		err = ipfsCmd.Run()
		if err != nil {
			panic(fmt.Errorf("error trying to remove bootstrap adresses: %v", err))
		}

		// Load sensormesh configurations to Viper
		shared.LoadConfigurationFromFile()

		// Set the node's initial configurati
		shared.ViperConfs.Set("name", nodename)
		shared.ViperConfs.Set("logfile", logfile)
		shared.ViperConfs.WriteConfig()

		fmt.Println("New sensormesh node " + shared.ViperConfs.GetString("name") + " created !")
	},
}

func init() {
	initCmd.Flags().StringVar(&swarmKey, "swarmKey", "", "IPFS private network swarm key, if none provided, creates a new one")
	initCmd.Flags().StringVar(&nodename, "nodename", "SensorMeshNode", "IPFS private network swarm key, if none provided, creates a new one")
	initCmd.Flags().StringVar(&logfile, "logfile", "~/.sensormesh/sensormesh.log", "Path destination for logfile, Defaults to '~/.sensormesh/sensormesh.log'")
	rootCmd.AddCommand(initCmd)
}
