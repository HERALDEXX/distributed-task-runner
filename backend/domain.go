package main

import (
	"time"
)

type JobStatus string

const (
	Pending   JobStatus = "pending"
	Running   JobStatus = "running"
	Completed JobStatus = "completed"
	Failed    JobStatus = "failed"
)

type Job struct {
	ID        string    `json:"id"`
	Payload   string    `json:"payload"`
	Status    JobStatus `json:"status"`
	Attempts  int       `json:"attempts"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Output    string    `json:"output"`
}

type Queue struct {
	Name string
	Jobs chan *Job
}

var defaultQueue = &Queue{
	Name: "default",
	Jobs: make(chan *Job, 100),
}

var jobStore = make(map[string]*Job)
