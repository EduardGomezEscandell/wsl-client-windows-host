package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: server.exe <address>")
		os.Exit(2)
	}

	url := os.Args[1]

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received message from %s", r.Host)
		fmt.Fprint(w, "Hello, World!")
	})

	log.Printf("Serving requests on %s\n", url)

	if err := http.ListenAndServe(url, nil); err != nil {
		log.Fatalf("Error: %v", err)
		os.Exit(1)
	}
}
