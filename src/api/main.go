package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	var e, exists = os.LookupEnv("FOO")
	if !exists {
		e = "no env"
	}
	var resp = fmt.Sprintf("Hello Sample GO API. foo: %s", e)
	w.Write([]byte(resp))
}

func main() {
	http.HandleFunc("/", handleRoot)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
