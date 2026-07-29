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

	for _, m := range missions {
		status := "Completed"
		if !m.Completed {
			status = "Incomplete"
		}
		fmt.Printf("Mission ID: %d, Title: %s, Status: %s\n", m.ID, m.Title, status)
	}
}