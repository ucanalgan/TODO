package main

import "fmt"

type Mission struct {
	ID int
	Title string
	Completed bool
}
func main(){
	missions := []Mission{}
	missions = append(missions, Mission{ID: 1, Title: "Mission 1", Completed: false})
	missions = append(missions, Mission{ID: 2, Title: "Mission 2", Completed: true})
	missions = append(missions, Mission{ID: 3, Title: "Mission 3", Completed: false})



	for {
		fmt.Println("1. Add Mission")
		fmt.Println("2. List Missions")
		fmt.Println("3. Complete Mission")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var title string
			fmt.Print("Enter mission title: ")
			fmt.Scanln(&title)
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
			var id int
			fmt.Print("Enter mission ID to complete: ")
			fmt.Scanln(&id)
			for i := range missions {
				if missions[i].ID == id {
					missions[i].Completed = true
					fmt.Println("Mission marked as completed!")
					break
				}
			}
		case 4:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}