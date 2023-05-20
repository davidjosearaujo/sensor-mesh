/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>
*/
package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"firemesh/cmd/shared"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/spf13/cobra"
)

var (
	IpfsApi				*shell.Shell
	IpfsPath			string
	swarmKey			string
	swarmKeyFilePath	= filepath.Join(shared.GetUserHomeDir(), ".ipfs", "swarm.key")
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

		var ipfsCmd *exec.Cmd

		// Searching for existing IPFS directory
		exists, _ := shared.Exists(filepath.Join(shared.GetUserHomeDir(), ".ipfs", "config"))
		if !exists {
			// Find IPFS binary absolute path
			_temp, err := exec.LookPath("ipfs")
			IpfsPath = _temp
			if err == nil {
				IpfsPath, err = filepath.Abs(IpfsPath)
			}
			if err != nil {
				fmt.Println("Error while looking for existing IPFS binary: ", err)
				os.Exit(1)
			}

			// Initiate IPFS service
			ipfsCmd = exec.Command(IpfsPath, "init")
			err = ipfsCmd.Run()
			if err != nil {
				fmt.Println("Error while creating IPFS config files: ", err)
				os.Exit(1)
			}
		} else {
			fmt.Println("IPFS config file already exists")
		}

		// Creates new swarm key file
		createSwarmKeyFile()

		// Clear bootstrap addresses
		IpfsApi = shell.NewShell("localhost:5001")
		_, err := IpfsApi.BootstrapRmAll()
		if err != nil {
			fmt.Println("Error while removing bootstrap addresses: ", err)
		}

		// Changing to DHT type routing
		ipfsCmd = exec.Command("ipfs", "config", "Routing.Type", "dht")
		err = ipfsCmd.Run()
		if err != nil {
			fmt.Println("Error while starting IPFS config files: ", err)
			os.Exit(1)
		}

		// Load firemesh configurations to Viper
		shared.LoadConfigurationFromFile()
	},
}

func init() {
	initCmd.Flags().StringVar(&swarmKey, "swarmKey", "", "IPFS private network swarm key, if none provided, creates a new one")

	rootCmd.AddCommand(initCmd)
}
