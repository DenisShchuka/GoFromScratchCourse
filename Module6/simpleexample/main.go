package main

import (
	"fmt"
	"time"
)

var done chan struct{}

func main() {
	done = make(chan struct{})
	go calc_f()
	for i := 'a'; i <= 'z'; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%c", i)
	}
	<-done
}

func calc_f() {
	for i := 0; i < 100; i++ {
		fmt.Print(i)
		time.Sleep(100 * time.Millisecond)
	}
	done <- struct{}{}
}
