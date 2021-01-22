package main

import (
	"fmt"
	"os"
)

func main() {
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	path := os.Getenv("PATH")

	fmt.Println("Path: " + path)
	fmt.Println("Go Root: " + goRoot)
	fmt.Println("Go Path: " + goPath)
}
