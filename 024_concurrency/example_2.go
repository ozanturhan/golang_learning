package main

import (
	"runtime"
	"time"
)

func main() {
	// Goroutine
	/*
		go xFunc()
		time.Sleep(100 * time.Millisecond)
		go islem()
		...
	*/

	runtime.GOMAXPROCS(1)
	go xFunc()
	time.Sleep(100 * time.Millisecond)
}

func xFunc() {
	for l := byte('a'); l <= byte('z'); l++ {
		go println(string(l))
	}
}
