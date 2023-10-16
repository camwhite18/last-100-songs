package main

import (
	"fmt"
	"log"
)

func main() {
	var stationURL, clientID, clientSecret string

	fmt.Println("Enter radio station URL: ")
	_, err := fmt.Scanln(&stationURL)
	if err != nil {
		log.Fatalf("Error reading station URL: %v", err)
	}

	fmt.Println("Enter Spotify client ID: ")
	_, err = fmt.Scanln(&clientID)
	if err != nil {
		log.Fatalf("Error reading client ID: %v", err)
	}

	fmt.Println("Enter Spotify client secret: ")
	_, err = fmt.Scanln(&clientSecret)
	if err != nil {
		log.Fatalf("Error reading client secret: %v", err)
	}

	se :=
}
