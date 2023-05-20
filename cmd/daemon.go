/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	orbitdb "berty.tech/go-orbit-db"
	client "github.com/ipfs/go-ipfs-http-client"
	"github.com/spf13/cobra"
)

var (
	orbit	*orbitdb.OrbitDB
)

func daemonLoop(){

}

// daemonCmd represents the daemon command
var daemonCmd = &cobra.Command{
	Use:   "daemon",
	Short: "Run a OrbitDB sensor logger",
	Long: `'firemesh daemon' runs a persistent firemesh daemon that can
query specified sensor and log their responses to a OrbitDB
log file, that will be shared between node in a same IPFS
private network.

The daemon will start by first configuring the current
machine as a node in a private IPFS network, and then
initialize IPFS's daemon`,
	Run: func(cmd *cobra.Command, args []string) {
		
		// Initiates IPFS daemon with pubsub enabled
		ipfsCmd := exec.Command(IpfsPath, "daemon", "--enable-pubsub-experiment")
		err := ipfsCmd.Run()
		if err != nil {
			panic(fmt.Errorf("Error while creating IPFS config files: %s", err))
		}

		c, err := client.NewLocalApi()
		if err != nil {
			panic(fmt.Errorf("failed to connect to local api: %s", err))
		}

		fmt.Println(c)

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		db, err := orbitdb.NewOrbitDB(ctx, c, &orbitdb.NewOrbitDBOptions{})
		if err != nil {
			panic(fmt.Errorf("failed to create new orbitdb: %s", err))
		}

		dbStore, err := db.Create(ctx, "test", "keyvalue", &orbitdb.CreateDBOptions{})
		if err != nil {
			panic(fmt.Errorf("failed to create new db store: %s", err))
		}

		fmt.Printf("dbStore address: %s\n", dbStore.Address())
		},
}

func init() {
	rootCmd.AddCommand(daemonCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daemonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daemonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
