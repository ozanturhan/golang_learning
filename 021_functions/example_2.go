package main

import "fmt"

// Functions: Multiple return value

func main() {
	// swap
	a, b := swap("Ozan", "Turhan")

	fmt.Println(a, b)

	numTerms, sum := addMultiple(3, 4, 5)

	fmt.Println("Added: ", numTerms, "Total:", sum)

	numTerms2, sum2 := addMultipleReturn(3, 4, 5)

	fmt.Println("Added: ", numTerms2, "Total:", sum2)
}

func swap(name, surname string) (string, string) {
	return surname, name
}

func addMultiple(terms ...int) (int, int) {
	result := 0
	for _, term := range terms {
		result += term
	}

	return len(terms), result
}


func addMultipleReturn(terms ...int) (length, result int) { // length and result defined in this place
	result = 0
	for _, term := range terms {
		result += term
	}

	length = len(terms)
	return
}