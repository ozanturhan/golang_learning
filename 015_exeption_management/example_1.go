package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// Exception Management

	/*
		file, _ := os.Open("file.txt")

		fmt.Println(file.Name())
	 */

	/*
		file, err := os.Open("file.txt")

		if err != nil {
			fmt.Println(err.Error)
		}

		fmt.Println(file.Name())
	 */

	/*
		i := -2
		if i < 0 {
			return 0, fmt.Errorf("matematik: çok kötü bir hata %g", i)
		}
	*/

	myError := errors.New("Bu bir hata")

	fmt.Println(myError)

	_, err := os.Open("abc.txt")

	checkError(err)
}

func checkError(err error) {
	if err != nil {

			fmt.Println("ERROR: ", err.Error())
			// logging
			os.Exit(1)

		// log.Fatal("ERROR: ", err)
	}
}