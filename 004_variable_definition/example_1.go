package main

import "fmt"

func main() {
	/*
		var message string
		message = "Hello Go"

		var message string = "Hello Go"

		var message = "Hello Go"

		var a, b, c int // a, b ve c 0
		var a, b, c int = 3, 4, 5
		var a, b, c = 3, 4, 5

		var a, b, c = 3, true, 4.5

		x := 55
		a, b := "abc", true

		a := "abc"
		b := 'a' // just one char with "'"
		c := `abc`
	*/

	a := "abc"
	b := 'a' // just one char with "'"
	c := `abc`

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
