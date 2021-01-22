package main

import "fmt"

func main() {

	// array
	a := [...]string{"a", "b", "d"}

	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	for i, v := range a {
		fmt.Println("Array item", i, "is", v)
	}

	// map
	cities := map[string]string{"Turkey": "Istanbul", "France": "Paris", "Italy": "Roma"}

	for key := range cities {
		fmt.Println(key, cities[key])
	}

	for key, val := range cities {
		fmt.Println(key, val)
	}
}
