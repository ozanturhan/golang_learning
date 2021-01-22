package main

import "fmt"

func main() {
	// basic array

	myArray1 := [3]int{}

	myArray1[0] = 32
	myArray1[1] = 23
	myArray1[2] = 54

	fmt.Println(myArray1)

	// color array

	var colors [3]string

	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println(colors)
	fmt.Println(colors[1])
	colors[1] = "White"
	fmt.Println(colors[1])

	var numbers = [5]int{4, 6, 7, 3, 66}
	fmt.Println(numbers)
	fmt.Println("Number of numbers", len(numbers))

	/// myArray2 := [4]int{4, 3, 5, 6}
	myArray2 := [...]int{4, 3, 5, 6, 7, 332}
	fmt.Println(myArray2)

	/*
		invalid array index 0
		myArray3 := [...]int{}
		myArray3[0] = 1
	*/

	var cars [3]string
	cars[0] = "Tesla"
	cars[1] = "Mercedes"
	cars[2] = "Jaguar"

	fmt.Println(len(cars))
	fmt.Println(cap(cars))

	i:=0
	for i<= len(cars) - 1 {
		fmt.Println(cars[i])
		i++
	}

	for i := 0; i < len(cars); i ++ {
		fmt.Println(cars[i])
	}

	for i, v := range cars {
		fmt.Println("index = ", i, "value = ", v)
	}
	for i := range cars {
		fmt.Println("index = ", i, "value = ", cars[i])
	}
}
