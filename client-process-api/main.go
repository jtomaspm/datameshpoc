package main

import "datamesh.poc/client-process-api/server"

func main() {
	s := server.New(&server.Config{
		Host:    "0.0.0.0",
		Port:    "8080",
		Context: "api/client",
	})
	s.Run()
}
