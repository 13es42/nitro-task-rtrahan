package main

import (
	"log"
	"nitro-task-rtrahan/api"
)

func main() {
	log.Println("Starting API server...")
	api.StartServer("8080")
}
