package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"
)

var n = flag.Int("n", 30e5, "Number of goroutines to create")

var ch = make(chan byte)
var counter = 0

func f() {
	counter++
	time.Sleep(100 * time.Millisecond)
	<-ch // Block this goroutine
}

func main() {
	flag.Parse()
	if *n <= 0 {
		fmt.Fprintf(os.Stderr, "invalid number of goroutines")
		os.Exit(1)
	}

	// Limit the number of spare OS threads to just 1
	runtime.GOMAXPROCS(1)

	// Make a copy of MemStats
	var m0 runtime.MemStats
	runtime.ReadMemStats(&m0)

	t0 := time.Now().UnixNano()
	for i := 0; i < *n; i++ {
		go f()
	}
	t1 := time.Now().UnixNano()

	// Make a copy of MemStats
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	fmt.Printf("Number of goroutines: %d\n", *n)
	fmt.Printf("Per goroutine:\n")
	fmt.Printf("  Memory: %.2f bytes\n", float64(m1.Sys-m0.Sys)/float64(*n))
	fmt.Printf("  Time:   %f ms\n", float64(t1-t0)/float64(*n)/1e6)
}
