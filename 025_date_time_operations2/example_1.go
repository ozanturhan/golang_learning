package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Now: %v\n", time.Now().Unix())
	time.Sleep(2 * time.Second)
	fmt.Printf("Now: %v\n", time.Now().Unix())
}
