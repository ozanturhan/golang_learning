package main

import "fmt"

var isConnected = false

func main() {
	fmt.Printf("Connection open %v\n", isConnected)
	databaseProcessing()
	fmt.Printf("Connection open %v\n", isConnected)
}

func databaseProcessing()  {
	connect()
	fmt.Println("Deferring disconnect")
	defer disconnect()
	fmt.Printf("Connection open: %v\n", isConnected)
	fmt.Println("Doing something")
}
func connect() {
	isConnected = true
	fmt.Println("Connected to database")
}

func disconnect() {
	isConnected = false

	fmt.Println("Disconnected!")
}