package main

import "fmt"

func main() {
	// IF

	a := 10
	b := 10

	if b > a {
		fmt.Println("Bigger")
	} else if b == a {
		fmt.Println("Equal")
	} else {
		fmt.Printf("Smaller")
	}

	/*
	foo := 1

	if foo == 1 {
		println("bar")
	}
	*/

	if foo := 2; foo == 1 {
		println("bar")
	} else {
		println("buz")
	}

	// fmt.Println(foo) unreachable
}
