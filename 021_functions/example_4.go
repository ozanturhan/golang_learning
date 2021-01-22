package main

// Anonymous Functions
func main() {
	addFunc := func(terms ...int) (length, sum int){
		for _, term := range terms {
			sum += term
		}
		length = len(terms)
		return
	}

	length, sum := addFunc(2,5, 5, 6, 7, 7, 3, 1)
	println(length, sum)
}

