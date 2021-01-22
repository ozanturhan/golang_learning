package main

import "fmt"

func main() {
	// array slice

	primes := [6]int{2, 3, 5, 7 ,11 , 13}

	var s []int = primes[1:4]
	fmt.Println(s)

	// array append

	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)
	// colors = append(colors[1:len(colors)], "White")
	colors = append(colors[1:], "White")
	fmt.Println(colors)
	colors = append(colors[1:len(colors)-2], "White")
	fmt.Println(colors)

	// make
	numbers := make([]int, 5, 5)

	numbers[0] = 123
	numbers[1] = 124
	numbers[2] = 125
	numbers[3] = 126
	numbers[4] = 127

	fmt.Println(numbers)

	numbers = append(numbers, 223)
	numbers = append(numbers, 224)
	numbers = append(numbers, 225)
	numbers = append(numbers, 227)
	numbers = append(numbers, 228)
	numbers = append(numbers, 228) // cap 5 to 10

	fmt.Println(numbers)
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
}
