package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"time"

	"github.com/google/uuid"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Distributed Task Runner!")
}

func enqueueHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Payload string `json:"payload"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	job := &Job{
		ID:        uuid.NewString(),
		Payload:   req.Payload,
		Status:    Pending,
		Attempts:  0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	jobStore[job.ID] = job
	defaultQueue.Jobs <- job

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(struct {
		ID     string    `json:"id"`
		Status JobStatus `json:"status"`
	}{
		ID:     job.ID,
		Status: job.Status,
	})
}

func getJobHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/jobs/"):]

	job, ok := jobStore[id]
	if !ok {
		http.Error(w, "Job not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(job)
}

func listJobsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	statusFilter := r.URL.Query().Get("status")

	jobs := []*Job{}
	for _, job := range jobStore {
		if statusFilter == "" || string(job.Status) == statusFilter {
			jobs = append(jobs, job)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

func clearJobsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Clear all jobs in memory
	jobStore = make(map[string]*Job)

	// Save empty jobStore to disk
	if err := saveJobsToDisk(); err != nil {
		http.Error(w, "failed to clear jobs", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
