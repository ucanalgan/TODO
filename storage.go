package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func saveMissionsToFile(missions []Mission) {
	data, err := json.MarshalIndent(missions, "", "  ")
	if err != nil {
		fmt.Println("Error encoding missions:", err)
		return
	}
	if err := os.WriteFile("missions.json", data, 0644); err != nil {
		fmt.Println("Error writing missions to file:", err)
	}
}

func loadMissionsFromFile() []Mission {
	data, err := os.ReadFile("missions.json")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("missions.json file not found. Starting with an empty mission list.")
		} else {
			fmt.Println("Error reading missions from file:", err)
		}
		return []Mission{}
	}
	var missions []Mission
	if err := json.Unmarshal(data, &missions); err != nil {
		fmt.Println("Error decoding missions:", err)
		return []Mission{}
	}
	return missions
}
