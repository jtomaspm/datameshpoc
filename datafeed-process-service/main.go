package main

import (
	"log"
	"sync"
	"time"

	"datamesh.poc/datafeed-process-service/feeder"
)

func main() {
	log.Println("Starting datafeed-process-service in 15 seconds...")
	time.Sleep(15 * time.Second)
	feeders := []feeder.Feeder{
		feeder.NewPersonFeeder("http://person-system-api:8080/api/person"),
	}
	wg := sync.WaitGroup{}
	wg.Add(len(feeders))
	for _, f := range feeders {
		go f.Feed(&wg)
	}
	wg.Wait()
}
