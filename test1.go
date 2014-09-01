package main

import (
	"fmt"
	"time"
)

func f(msg string) {
	fmt.Println(msg)
}

func main() {
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(1000)
}
