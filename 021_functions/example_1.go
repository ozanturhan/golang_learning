package main

import "fmt"

func main() {

	name := "Ozan"
	sayHello(name)
	println(name)

	namePointer := "Ozan"
	sayHelloWithPointer(&namePointer)
	println(namePointer)

	fmt.Println(add(3, 4, ""))
}

func sayHello(name string) {
	println("Hello", name)
	name = "Zeynep"
}

func sayHelloWithPointer(name *string) {
	println("Hello", *name)
	*name = "Zeynep"
}

func add(x, y int, z string) int {
	return x + y
}
