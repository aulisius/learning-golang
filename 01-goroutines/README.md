# GoRoutine Basics

## What are GoRoutines?

- Goroutines are functions which executes concurrently along with the main program flow.
- Lightweight, independently executing concurrent unit of work
- Unlike traditional operating system threads,Goroutines are multiplexed to a fewer number of OS threads, managed by the Go runtime and have
  a much smaller memory footprint. Go allows you to create thousands or even millions of goroutines within a program without significant
  performance degradation.
- Created using the <b>go</b> keyword followed by a function call or function literal

## Memory Allocation for GoRoutines

- In Go, memory for goroutines is allocated on the heap.
- When a goroutine is created, the Go runtime sets up a stack for that goroutine. The goroutine stack is used to store local variables,
  function calls, and other execution-related data. It is a fixed-size segment of memory that grows and shrinks as needed during the
  execution of the goroutine.
- The Go runtime handles the management of goroutine stacks, allowing efficient and scalable concurrent programming.

## Example Walkthrough

- Here in the example, we are calling function <b>say</b> in the separate goroutine and main will be executing concurrently along with it in
  any order.