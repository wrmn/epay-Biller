package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// ChannelKafka started
	log.Println("Service Started!")

	// Setting up log file
	// set permission to read/write log file
	// read/write to existing log file, if there is none it will create new log file
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Found error in log ", err)
	}
	log.SetOutput(file)

	// Setting up HTTP Listener and Handler
	// router will handle any request at any endpoint available in server()
	path := pathHandler()
	// listen to specific address and handler
	address := "localhost:6001"
	err = http.ListenAndServe(address, path)
	log.Println("Server started at", address)
	if err != nil {
		log.Fatal(err.Error())
	}
}
