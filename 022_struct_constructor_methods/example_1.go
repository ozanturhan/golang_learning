package main

import "fmt"

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

// constructor
func NewHuman() *Human {
	h := new(Human)
	return h
}

// constructor with values
func NewHumanWithValues(firstName, lastName string, age int) *Human {
	h := Human{FirstName: firstName, LastName: lastName, Age: age}
	return &h
}

/*
func NewHumanWithValues(firstName, lastName string) (h *Human) {
	h = &Human{FirstName: firstName, LastName: lastName}
	return
}
*/

func main() {

	// v1
	x := Human{FirstName: "Ozan"}

	fmt.Println(x.FirstName)

	// v2
	y := new(Human)
	y.FirstName = "Zeynep"

	fmt.Println(y.FirstName)

	// v3
	z := NewHuman()
	z.FirstName = "Abdullah"
	z.Age = 3

	fmt.Println(z.Age)

	// v4
	a := NewHumanWithValues("Ozan", "Turhan", 29)

	fmt.Println(a.FirstName, a.LastName, a.Age)
}
