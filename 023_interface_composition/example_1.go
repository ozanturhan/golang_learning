package main

import "fmt"

func CarExecute(c Carface) {
	fmt.Println('\n')
	fmt.Println("Car Info: \n" + c.Information())

	msg := ""

	isRun := c.Run()
	if isRun {
		msg = "is run"
	} else {
		msg = "is not run"
	}

	fmt.Println("Car " + msg + ".")

	isStop := c.Stop()
	if isStop {
		msg = "is stop"
	} else {
		msg = "is not stop"
	}

	fmt.Println("Car " + msg + ".")

}

func main() {
	ferrari := new(Ferrari)
	ferrari.Brand = "Ferrari"
	ferrari.Model = "F50"
	ferrari.Color = "Red"
	ferrari.Speed = 340
	ferrari.Price = 1.4
	ferrari.Special = true
	fmt.Println(ferrari.Information())

	CarExecute(ferrari)
}

// Interface

type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}

// Base structs
type Car struct {
	Brand string
	Model string
	Color string
	Speed int
	Price float64
}

type SpecialProduction struct {
	Special bool
}

// Ferrari embed type
type Ferrari struct {
	Car // composition
	SpecialProduction
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
	return true
}

func (ferrari *Ferrari) Information() string {
	ret := "\t" + ferrari.Brand + " " + ferrari.Model + "\n\t" + "Color: " + ferrari.Color
	add := "Yes"

	if ferrari.Special {
		ret += "\n\t" + "Special: " + add
	}

	return ret
}

// Lamborghini embed type
type Lamborghini struct {
	Car // composition
	SpecialProduction
}

func (_ Lamborghini) Run() bool {
	return true
}

func (_ Lamborghini) Stop() bool {
	return true
}

func (lamborghini *Lamborghini) Information() string {
	ret := "\t" + lamborghini.Brand + " " + lamborghini.Model + "\n\t" + "Color: " + lamborghini.Color
	add := "Yes"

	if lamborghini.Special {
		ret += "\n\t" + "Special: " + add
	}

	return ret
}

// Mercedes embed type
type Mercedes struct {
	Car // composition
}

func (_ Mercedes) Run() bool {
	return true
}

func (_ Mercedes) Stop() bool {
	return true
}

func (mercedes *Mercedes) Information() string {
	ret := "\t" + mercedes.Brand + " " + mercedes.Model + "\n\t" + "Color: " + mercedes.Color
	return ret
}
