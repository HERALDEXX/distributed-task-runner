package main

import (
	"fmt"
	"net/http"
)

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := loadJobsFromDisk(); err != nil {
		fmt.Println("Failed to load jobs from disk:", err)
	}

	numWorkers := 3
	for i := 1; i <= numWorkers; i++ {
		w := &Worker{
			ID:    i,
			Queue: defaultQueue,
			Quit:  make(chan struct{}),
		}
		w.Start()
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", helloHandler)
	
	mux.HandleFunc("/jobs", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			enqueueHandler(w, r)
		case http.MethodGet:
			listJobsHandler(w, r)
		case http.MethodDelete:
			clearJobsHandler(w, r)	
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/jobs/", getJobHandler)
	mux.HandleFunc("/jobs/clear", clearJobsHandler)

	fmt.Println("Server listening on :8080")
	if err := http.ListenAndServe(":8080", withCORS(mux)); err != nil {
		panic(err)
	}
}
