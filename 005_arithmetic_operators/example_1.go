package main

import "fmt"

func main() {
	// Operators

	a := 10
	b := 20

	total := a + b

	fmt.Println("total:", total)

	total *= 20

	fmt.Println("total:", total)

	total = 10 / 2
	fmt.Println("total:", total)

	total = a % b
	fmt.Println("total:", total)

	total++
	fmt.Println("total:", total)

	total--
	fmt.Println("total:", total)
}
