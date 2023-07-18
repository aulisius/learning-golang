package main

import (
	"fmt"
	"time"
)

func acceptNumbers(ch, quit chan int) {

	for i := 0; i < 10; i++ {
	    fmt.Println("Producing =>", i)
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	quit <- 0
}

func printNumbers(ch, quit chan int) {

	for {
		select {
		case value := <-ch:
			fmt.Println("Consuming =>", value)
			fmt.Println("-----------------")
		case <-quit:
			fmt.Println("Quit")
			return
		}
	}
}

func main() {

	ch := make(chan int)
	quit := make(chan int)
	go acceptNumbers(ch, quit)
	printNumbers(ch, quit)
}
