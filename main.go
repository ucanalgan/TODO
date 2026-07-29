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



func main() {
	missions := []Mission{}
	missions = append(missions, Mission{ID: 1, Title: "Mission 1", Completed: false})
	missions = append(missions, Mission{ID: 2, Title: "Mission 2", Completed: true})
	missions = append(missions, Mission{ID: 3, Title: "Mission 3", Completed: false})



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
			for i := range missions {
				if missions[i].ID == id {
					missions[i].Completed = true
					fmt.Println("Mission marked as completed!")
					found = true
					break
				}
			}
			if !found {
				fmt.Println("Invalid mission ID")
			}
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}