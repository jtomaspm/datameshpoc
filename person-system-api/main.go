package main

import (
	"datamesh.poc/person-system-api/server"
)

func main() {
	s := server.New(&server.Config{
		Host:    "0.0.0.0",
		Port:    "8080",
		Context: "api/person",
	})
	s.Run()
}
