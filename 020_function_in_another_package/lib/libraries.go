package main

import (
	"../utils"
	"fmt"
)

func main()  {

	n1, l1 := utils.FullName("Ozan", "Turhan")
	fmt.Printf("Fullname: %v number of chars: %v\n\n", n1, l1)

	n2, l2 := utils.FullNameNakedReturn("Ozan", "Turhan")
	fmt.Printf("Fullname: %v number of chars: %v\n\n", n2, l2)
}
