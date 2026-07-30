package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type Mission struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func readLine(prompt string) string {
	fmt.Print(prompt)
	line, _ := reader.ReadString('\n')
	return strings.TrimSpace(line)
}
func readInt(prompt string) (int, error) {
	for {
		input := readLine(prompt)
		value, err := strconv.Atoi(input)
		if err == nil {
			return value, nil
		}
		fmt.Println("Invalid input. Please enter a valid integer.")
	}
}
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

func main() {
	missions := loadMissionsFromFile()
	nextID := 1
	for _, m := range missions {
		if m.ID >= nextID {
			nextID = m.ID + 1
		}
	}

	for {
		fmt.Println("1. Add Mission")
		fmt.Println("2. List Missions")
		fmt.Println("3. Complete Mission")
		fmt.Println("4. Delete Mission")
		fmt.Println("5. Edit Mission")
		fmt.Println("6. Exit")

		choice, err := readInt("Enter your choice: ")
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}

		switch choice {
		case 1:
			title := readLine("Enter mission title: ")
			newMission := Mission{
				ID:        nextID,
				Title:     title,
				Completed: false,
			}
			if title == "" {
				fmt.Println("Mission title cannot be empty!")
			} else {
				missions = append(missions, newMission)
				nextID++
				saveMissionsToFile(missions)
				fmt.Println("Mission added successfully!")
			}
		case 2:
			fmt.Println("Missions:")
			for _, m := range missions {
				status := "Completed"
				if !m.Completed {
					status = "Incomplete"
				}
				fmt.Printf("Mission ID: %d, Title: %s, Status: %s\n", m.ID, m.Title, status)
			}
		case 3:
			id, err := readInt("Enter mission ID to complete: ")
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid integer.")
				continue
			}
			found := false
			changed := false
			for i := range missions {
				if missions[i].ID == id {
					found = true
					if missions[i].Completed {
						fmt.Println("Mission is already completed!")
					} else {
						missions[i].Completed = true
						fmt.Println("Mission marked as completed!")
						changed = true
					}
					break
				}
			}
			if !found {
				fmt.Println("Invalid mission ID")
			} else if changed {
				saveMissionsToFile(missions)
			}
		case 4:
			id, err := readInt("Enter mission ID to delete: ")
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid integer.")
				continue
			}
			found := false
			for i := range missions {
				if missions[i].ID == id {
					found = true
					missions = append(missions[:i], missions[i+1:]...)
					fmt.Println("Mission deleted successfully!")
					break
				}
			}
			if !found {
				fmt.Println("Invalid mission ID")
			} else {
				saveMissionsToFile(missions)
			}
		case 5:
			id, err := readInt("Enter mission ID to edit: ")
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid integer.")
				continue
			}
			found := false
			changed := false
			for i := range missions {
				if missions[i].ID == id {
					found = true
					title := readLine("Enter new mission title: ")
					if title == "" {
						fmt.Println("Mission title cannot be empty!")
					} else {
						missions[i].Title = title
						changed = true
						fmt.Println("Mission edited successfully!")
					}
					break
				}
			}
			if !found {
				fmt.Println("Invalid mission ID")
			} else if changed {
				saveMissionsToFile(missions)
			}
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}
