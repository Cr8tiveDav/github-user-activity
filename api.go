package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func fetchEvents(username string) (Events, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(url)
	if err != nil {
		return Events{}, err
	}
	defer resp.Body.Close()

	type ErrorResponse struct {
		Message     string `json:"message"`
		DocumentURL string `json:"documentation_url"`
		StatusCode  int    `json:"status_code"`
	}

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		var errorResp ErrorResponse
		json.Unmarshal(body, &errorResp)
		fmt.Fprintf(os.Stderr, "Error: API request failed with status %d %s\n", resp.StatusCode, errorResp.Message )
		os.Exit(1)
	}

	var events Events
	err = json.NewDecoder(resp.Body).Decode(&events)
	if err != nil {
		return Events{}, err
	}

	return events, nil
}
