package main

import (
	"log"
	"net/http"
	"time"
	"os"
	"strings"
)

func main() {
	if 2 == len(os.Args) {
		// Get url from command line args
		url := os.Args[1]

		log.Println("URL set to: " + url)

		// Retrieve url data
		log.Println("Retrieving url data...")
		client := http.Client{
			Timeout: time.Duration(5 * time.Second),
		}

		res, err := client.Head(url)
		if err != nil {
			log.Fatal(err)
		}

		// Print status code
		log.Printf("Status code:\t%d\n", res.StatusCode)

		// Print headers
		log.Println("Start printing headers...")
		for key, value := range res.Header {
			log.Println("\t" + key + ": " + strings.Join(value, ","))
		}
	} else {
		log.Printf("Usage: %s <url>\n", os.Args[0])
	}
}
