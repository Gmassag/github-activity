package main

import (
	"fmt"
	"os"

	"github-activity/internal/storage"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		return
	}
	username := os.Args[1]

	events, err := storage.FetchUserEvents(username)
	if err != nil {
		fmt.Println("Errore:", err)
		return
	}

	for _, e := range events {
		switch e.Type {
		case "PushEvent":
			fmt.Printf("- %s pushed %d commits to %s\n", e.Actor.Login, len(e.Payload.Commits), e.Repo.Name)
		case "WatchEvent":
			fmt.Printf("- %s starred %s\n", e.Actor.Login, e.Repo.Name)
		case "IssuesEvent":
			fmt.Printf("- %s opened an issue in %s\n", e.Actor.Login, e.Repo.Name)
		default:
			fmt.Printf("- %s did %s in %s\n", e.Actor.Login, e.Type, e.Repo.Name)
		}
	}
}
