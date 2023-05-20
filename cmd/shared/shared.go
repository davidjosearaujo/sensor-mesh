/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>
*/
package shared

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

var (
	ViperConfs     = viper.New()
	ConfigFilePath = filepath.Join(GetUserHomeDir(), ".firemesh", "config.yaml")
)

func GetUserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("Error getting user home directory: %v\n", err)
		os.Exit(1)
	}
	return usr.HomeDir
}

func AddSensor(name string,
	port string,
	baud int,
	size int,
	parity string,
	stop string,
	interval int) {

	newSensor := map[string]interface{}{
		"name":     name,
		"baud":     baud,
		"parity":   parity,
		"port":     port,
		"size":     size,
		"stop":     stop,
		"interval": interval,
	}

	sensorList := ViperConfs.Get("sensors").([]interface{})
	sensorList = append(sensorList, newSensor)
	ViperConfs.Set("sensors", sensorList)
	err := ViperConfs.WriteConfig()
	if err != nil {
		fmt.Printf("Error updating config file: %v\n", err)
		os.Exit(1)
	}
}

func RemoveSensor(name string) {
	if ViperConfs.Get("sensors") == nil {
		fmt.Println("Error: There are no sensors to remove")
		os.Exit(1)
	}

	sensorList := ViperConfs.Get("sensors").([]interface{})
	updatedSensors := make([]interface{}, 0, len(sensorList))

	for _, sensor := range sensorList {
		sensorMap := sensor.(map[string]interface{})
		if sensorMap["name"].(string) != name {
			updatedSensors = append(updatedSensors, sensor)
		}
	}

	ViperConfs.Set("sensors", updatedSensors)
	err := ViperConfs.WriteConfig()
	if err != nil {
		fmt.Printf("Error updating config file: %v\n", err)
		os.Exit(1)
	}
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
			fmt.Printf("Error creating directories: %v\n", err)
			os.Exit(1)
		}
		file, err = os.Create(ConfigFilePath)
		if err != nil {
			fmt.Printf("Error generating config file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		file.Write(OriginalConfigFileContent)
	}

	path, name := filepath.Split(ConfigFilePath)
	ViperConfs.AddConfigPath(path)
	ViperConfs.SetConfigName(strings.TrimSuffix(name, filepath.Ext(name)))
	ViperConfs.SetConfigType("yaml")
	ViperConfs.ReadInConfig()
}

// Default content of configuration file
var OriginalConfigFileContent = []byte(`# configuration verification interval in seconds
refresh: 60

# List of sensors
sensors:

# Log files and respective locations
logFiles:
  public: "~/.firemesh/community.log"
  private: "~/.firemesh/private.log"
`)
