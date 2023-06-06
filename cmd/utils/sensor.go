/*
Copyright © 2023 David Araújo <davidaraujo98@github.io>
*/
package utils

import (
	"fmt"
)

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

	// Check if list is empty
	sensorList := []interface{}{}
	if ViperConfs.IsSet("sensors") {
		// Initialize "channels" key with an empty slice
		sensorList = ViperConfs.Get("sensors").([]interface{})
	}

	sensorList = append(sensorList, newSensor)
	ViperConfs.Set("sensors", sensorList)
	err := ViperConfs.WriteConfig()
	if err != nil {
		panic(fmt.Errorf("error adding new sensor: %v", err))
	}
}

func RemoveSensor(name string) {
	if ViperConfs.Get("sensors") == nil {
		panic(fmt.Errorf("there are no sensors to remove"))
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
		panic(fmt.Errorf("error removing sensor: %v", err))
	}
}
