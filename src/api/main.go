package main

import (
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello Sample Go API"))
}

func main() {
	http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
