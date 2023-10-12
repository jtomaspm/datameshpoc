package main

import (
	"sync"

	"datamesh.poc/log-process-api/worker"
)

func main() {
	wg := sync.WaitGroup{}
	workers := []worker.Worker{
		worker.NewSubscriber(),
	}
	wg.Add(len(workers))
	for _, w := range workers {
		go w.Run(&wg)
	}
	wg.Wait()
}
