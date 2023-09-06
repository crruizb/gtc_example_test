package main

import (
	"gtc_example/api"
)

func main() {
	println("Starting main...")
	server := api.CreateServer("8080")
	server.Run()
}
