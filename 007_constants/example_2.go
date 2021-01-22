package main

import "fmt"

type Brand string

const (
	FACEBOOK Brand = "Facebook"
	MICROSOFT Brand = "Microsoft"
	GOOGLE Brand = "Google"
)

func PrintBrand(brand Brand)  {
	fmt.Println(brand)
}

func main() {
	PrintBrand(FACEBOOK)
	PrintBrand(MICROSOFT)
	PrintBrand(GOOGLE)
}
