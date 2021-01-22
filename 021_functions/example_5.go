package main

import "fmt"

// defer

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	/*
	result:
	hello
	world
	 */
}
