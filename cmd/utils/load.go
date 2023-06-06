/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>
*/
package utils

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var (
	ViperConfs     = viper.New()
	ConfigFilePath = filepath.Join(GetUserHomeDir(), ".sensormesh", "config.yaml")
	LogFilePath    = filepath.Join(GetUserHomeDir(), ".sensormesh", "sensormesh.log")
)

func GetUserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(fmt.Errorf("error getting user home directory: %v", err))
	}
	return usr.HomeDir
}

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	return false, err
}

func LoadConfigurationFromFile() {
	var (
		file *os.File
		err  error
	)

	exists, _ := Exists(ConfigFilePath)
	if !exists {
		err = os.MkdirAll(filepath.Dir(ConfigFilePath), 0700)
		if err != nil {
			panic(fmt.Errorf("error creating directories: %v", err))
		}
		file, err = os.Create(ConfigFilePath)
		if err != nil {
			panic(fmt.Errorf("error generating config file: %v", err))
		}
		file.Write([]byte(""))
		file.Close()
		fmt.Println("[+] Configuration file created at " + ConfigFilePath)

		file, err = os.Create(LogFilePath)
		if err != nil {
			panic(fmt.Errorf("error generating log file: %v", err))
		}
		defer file.Close()
		file.Write([]byte(""))
		fmt.Println("[+] Log file created at " + LogFilePath)
	}

	path, name := filepath.Split(ConfigFilePath)
	ViperConfs.AddConfigPath(path)
	ViperConfs.SetConfigName(strings.TrimSuffix(name, filepath.Ext(name)))
	ViperConfs.SetConfigType("yaml")
	ViperConfs.ReadInConfig()
	ViperConfs.WatchConfig()
}

func LocalIPFSApiAddress() string {
	// Find local IPFS node API address
	out, err1 := exec.Command("ipfs", "config", "Addresses.API").Output()
	if err1 != nil {
		panic(err1)
	}
	ipfsApiIP := strings.Split(string(out), "/")[2] + ":" + strings.TrimSuffix(strings.Split(string(out), "/")[4], "\n")

	return ipfsApiIP
}
