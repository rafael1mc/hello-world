package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Handler)
	envPort := os.Getenv("API_PORT")
	if envPort == "" {
		envPort = "3000"
	}
	port := fmt.Sprintf(":%s", envPort)
	fmt.Println("Listening on port", port)
	http.ListenAndServe(port, nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bye, World!")
}
