package main

import (
	"fmt"
)

func main() {

	ch := make(chan int, 2) //channel of size 2
	ch <- 2
	ch <- 3 //it won't block as there is space in the buffer
	//ch <- 4
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//show deadlock case
