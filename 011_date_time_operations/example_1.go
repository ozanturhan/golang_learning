package main

import (
	"fmt"
	"time"
)

func main() {
	// Console: Date & Time Operations

	t := time.Date(2016, time.November, 10, 20, 0,0,0, time.UTC)

	fmt.Printf("Go launch at %s\n", t)

	fmt.Println("--------------------------------------------------------------------------")

	now := time.Now()

	fmt.Printf("Go launch at %s\n", now)

	fmt.Println("--------------------------------------------------------------------------")

	fmt.Println("The month is", t.Month())
	fmt.Println("The day is", t.Day())
	fmt.Println("The weekday is", t.Weekday())

	fmt.Println("--------------------------------------------------------------------------")

	tomorrow := now.AddDate(0, 0, 1)

	fmt.Printf(
		"Tomorrom is %v, %v, %v, %v\n",
		tomorrow.Weekday(),
		tomorrow.Month(),
		tomorrow.Day(),
		tomorrow.Year(),
		)

	longFormat := "Monday, January 2, 2006"
	fmt.Println("Tomorrow is", tomorrow.Format(longFormat))

	shortFormat := "1/2/06"
	fmt.Println("Tomorrow is", tomorrow.Format(shortFormat))
}
