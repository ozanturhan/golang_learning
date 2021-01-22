package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	print := fmt.Println

	xTime := time.Date(2021, 01, 11, 20, 30, 0, 0, time.UTC)
	now := time.Now()

	print(now.Year())
	print(now.Month())
	print(now.Hour())
	print(now.Minute())
	print(now.Second())
	print(now.Nanosecond())
	print(now.Location())

	print(now.Weekday())

	// Compare

	print(xTime.Before(now))
	print(xTime.After(now))

	print(xTime, now)
	diff := now.Sub(xTime)
	print(math.Floor(diff.Hours() / 24))
}
