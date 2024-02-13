package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: client <address>")
		os.Exit(2)
	}

	addr := os.Args[1]
	client := &http.Client{
		Timeout: 5 * time.Second, // set 5-second timeout
	}

	resp, err := client.Get(addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to %s: %v\n", addr, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	fmt.Printf("Connected to %s: %s\n", addr, resp.Status)
}
