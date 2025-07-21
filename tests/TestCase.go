package main

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	url         = "https://flight.asialuxe.dev/api/v1/charter-search"
	totalReq    = 200
	concurrency = 10
)

var jsonData = []byte(`{
    "hex":"5987c9ca9dff423ff02e04fd3071a7db",
    "depute":"TAS",
    "arrive":"AYT",
    "date_from":"13.08.2025",
    "date_to":null,
    "passAdult":1,
    "passChild":0,
    "passInfants":0,
    "cabin":"Y",
    "currency":"USD"
}`)

func sendRequest(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("❌ Request %d: Failed to create request: %s\n", id, err)
		return
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Request %d: Error: %s\n", id, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("✅ Request %d: Status Code: %d\n", id, resp)
}

func main() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency)

	for i := 1; i <= totalReq; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int) {
			defer func() { <-sem }()
			sendRequest(i, &wg)
		}(i)
	}

	wg.Wait()
	fmt.Println("✅ All requests completed.")
}
