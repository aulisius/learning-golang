package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("sending value to channel")
		ch <- 42
	}()

	val := <-ch
	fmt.Println("received value from channel:", val)
}
