package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

type Worker struct {
	ID    int
	Queue *Queue
	Quit  chan struct{}
}

func (w *Worker) Start() {
	fmt.Printf("Worker %d started\n", w.ID)
	go func() {
		for {
			select {
			case job := <-w.Queue.Jobs:
				fmt.Printf("Worker %d picked job %s\n", w.ID, job.ID)
				w.process(job)
			case <-w.Quit:
				fmt.Printf("Worker %d stopping\n", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) process(job *Job) {
	job.Status = Running
	job.Attempts++
	job.UpdatedAt = time.Now()
	fmt.Printf("Worker %d processing job %s (attempt %d): %s\n", w.ID, job.ID, job.Attempts, job.Payload)

	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", job.Payload)
	} else {
		cmd = exec.Command("sh", "-c", job.Payload)
	}

	out, err := cmd.CombinedOutput()

	if err != nil {
		job.Status = Failed
		job.UpdatedAt = time.Now()
		job.Output = fmt.Sprintf("error: %v\noutput: %s", err, string(out))

		fmt.Printf("Worker %d failed job %s ❌: %v\n", w.ID, job.ID, err)

		saveJobsToDisk()

		if job.Attempts < 3 {
			go func() {
				time.Sleep(1 * time.Second)
				defaultQueue.Jobs <- job
			}()
		}
		return
	}

	job.Status = Completed
	job.UpdatedAt = time.Now()
	job.Output = string(out)

	fmt.Printf("Worker %d completed job %s ✅\n", w.ID, job.ID)

	saveJobsToDisk()
}
