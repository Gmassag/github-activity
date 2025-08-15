package storage

import (
	"encoding/json"
	"fmt"
	"github-activity/internal/model"
	"net/http"
)

func FetchUserEvents(username string) ([]model.Event, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status non OK: %s", resp.Status)
	}

	var events []model.Event
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, err
	}

	return events, nil
}
