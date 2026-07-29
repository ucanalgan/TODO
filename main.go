package main

import ("bufio"
		"encoding/json"
		"fmt"
		"os"
		"strings"
		"strconv"
		)

var reader = bufio.NewReader(os.Stdin)

type Mission struct {
	ID        int		`json:"id"`
	Title     string 	`json:"title"`
	Completed bool 		`json:"completed"`
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
	if err:= os.WriteFile("missions.json", data, 0644); err != nil {
		fmt.Println("Error writing missions to file:", err)
	}
}

func loadMissionsFromFile() []Mission {
	data, err := os.ReadFile("missions.json")
	if err != nil {
		fmt.Println("Error reading missions from file:", err)
		return []Mission{}
	}
	var missions []Mission
	json.Unmarshal(data, &missions)
	return missions
}



func main() {
	missions := loadMissionsFromFile()

	for {
		fmt.Println("1. Add Mission")
		fmt.Println("2. List Missions")
		fmt.Println("3. Complete Mission")
		fmt.Println("4. Exit")
		
		choice,err := readInt("Enter your choice: ")
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer.")
			continue
		}


		switch choice {
		case 1:
			title := readLine("Enter mission title: ")
			newMission := Mission{
				ID: len(missions) + 1,
				Title: title,
				Completed: false,
			}
			missions = append(missions, newMission)
			saveMissionsToFile(missions)
			fmt.Println("Mission added successfully!")
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
			id,err := readInt("Enter mission ID to complete: ")
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
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}