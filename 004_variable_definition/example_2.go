package main

import "fmt"

var x = 3 // var is global keyword
// y := 4 unexpected

// const description - missing value

const description = "Hello Const"
const pi = 3.14

var a = "Abc"
var b, c = "Abc", "xyz"
var d string

var (
	variable1 = "x"
	variable2 = "y"
)

func main() {
	// d = "Hello global variable"
	// description = "Assign new value" - cannot assign const

	// var x = 5

	fmt.Println(variable1)
}