package main

import "fmt"

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
			completedCount := 0
			if len(missions) == 0 {
				fmt.Println("No missions found!")
				continue
			}
			fmt.Println("Missions:")
			for _, m := range missions {
				status := "Completed"
				if m.Completed {
					completedCount++
				} else {
					status = "Incomplete"
				}
				fmt.Printf("Mission ID: %d, Title: %s, Status: %s\n", m.ID, m.Title, status)
			}
			fmt.Printf("Total Missions: %d, Completed: %d, Incomplete: %d\n", len(missions), completedCount, len(missions)-completedCount)
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
