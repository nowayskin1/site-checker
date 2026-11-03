package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"https://github.com",
		"https://google.com",
		"https://example.com",
	}

	client := http.Client{Timeout: 5 * time.Second}

	for _, site := range sites {
		start := time.Now()
		resp, err := client.Get(site)
		if err != nil {
			fmt.Printf("DOWN  %s (%v)\n", site, err)
			continue
		}
		resp.Body.Close()
		fmt.Printf("UP    %s  status=%d  time=%v\n", site, resp.StatusCode, time.Since(start).Round(time.Millisecond))
	}
}
