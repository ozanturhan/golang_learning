package main

import (
	"fmt"
	"time"
)

// Date and Time Operations

func main() {
	seperater()

	t := time.Date(2018, time.November, 10, 20, 0, 0, 0, time.UTC)
	fmt.Printf("Date: %s\n", t)

	seperater()

	now := time.Now()
	fmt.Printf("Now: %s\n", now)

	seperater()

	fmt.Println("Month:", now.Month())
	fmt.Println("Day:", now.Day())
	fmt.Println("Weekday:", now.Weekday())

	seperater()

	tomorrow := now.AddDate(0, 0, 1)
	fmt.Printf("Tomorrow is %v, %v, %v, %v\n", tomorrow.Weekday(), tomorrow.Month(), tomorrow.Day(), tomorrow.Year())

	seperater()

	longFormat := "Monday, January 02, 2006"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))

	seperater()

	shortFormat := "2006/01/02"
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))

}

func seperater() {
	fmt.Println("---------------------------------")
}
