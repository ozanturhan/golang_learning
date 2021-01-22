package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	log "github.com/koding/logging"
)

func main() {
	// fmt
	// fmt.Println("This is an example")

	// fmt.Println("My favorite number is", rand.Intn(10))

	// strings
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("unit_test", "unit"))
	fmt.Println(strings.HasSuffix("file.rar", "rar"))
	fmt.Println(strings.Index("test", "e"))

	color.Red("Error Message")

	log.Info("Exit")
}
