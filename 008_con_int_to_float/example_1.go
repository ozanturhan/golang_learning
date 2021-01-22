package main

import "fmt"

func main() {
	// type conversions

	/*
		var myString = "1"
		var x = 10
		var f float32 = 2.0

		fmt.Println(myString, x, f)
		// convert string to int

		number, _ := strconv.Atoi(myString)
		fmt.Println("sonuc: " + strconv.Itoa(number))
		result := number + 2
		fmt.Println(result)
	 */

	var i int = 55
	// var f1 float64 = i - error
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Println(i, f, u)
}

