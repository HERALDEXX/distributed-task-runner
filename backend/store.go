package main

import (
	"encoding/json"
	"os"
)

func saveJobsToDisk() error {
	data, err := json.MarshalIndent(jobStore, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile("jobs.json", data, 0644)
}

func loadJobsFromDisk() error {
	if _, err := os.Stat("jobs.json"); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile("jobs.json")
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &jobStore)
}
