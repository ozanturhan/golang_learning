package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Console: Input

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter text: ")

	str, _ := reader.ReadString('\n')
	fmt.Println(str)

	fmt.Print("Enter a number: ")
	str, _ = reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Value of number:", f)
	}
}
