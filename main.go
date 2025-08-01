package main

import (
  "fmt"
  "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello, Distributed Scheduler!")
}

func main() {
  http.HandleFunc("/", helloHandler)
  fmt.Println("Server listening on :8080")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
