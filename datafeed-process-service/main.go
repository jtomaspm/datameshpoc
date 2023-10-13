package main

import "datamesh.poc/datafeed-process-service/server"

func main() {
	// log.Println("Starting datafeed-process-service in 15 seconds...")
	// time.Sleep(15 * time.Second)
	// feeders := []feeder.Feeder{
	// 	//feeder.NewPersonFeeder("http://person-system-api:8080/api/person"),
	// }
	// wg := sync.WaitGroup{}
	// wg.Add(len(feeders))
	// for _, f := range feeders {
	// 	go f.Feed(&wg)
	// }
	// wg.Wait()
	s := server.New(&server.Config{
		Host: "0.0.0.0",
		Port: "8080",
	})
	s.Run()
}
