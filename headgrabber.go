package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	url := "http://www.amazon.com/"
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	res, err := client.Head(url)
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range res.Header {
		log.Println(k+":", v)
	}
}
