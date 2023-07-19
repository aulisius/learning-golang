# Channels

## Basics

In Golang, a channel is a communication mechanism that allows two or more goroutines to communicate and synchronize their actions. A channel
is a way to pass messages between goroutines, allowing them to coordinate their work and share data without having to use explicit locks or
other synchronization mechanisms.

A channel is created using the `make()` function, which takes the type of values that the channel will hold as an argument.

For example, to create an unbuffered channel that holds integers, you would use the following code:

```go
ch := make(chan int)
```

Once a channel is created, you can send values into it using the <- operator and receive values from it using the same operator in a
different context. For example, to send the integer value 42 into a channel, you would use the following code:

```go
ch <- 42
```

And to receive a value from a channel, you would use the following code:

```go
x := <-ch
```

Channels are blocking. If a goroutine tries to read from a empty channel then it will be blocked and program execution keeps on waiting
until it recieves a value.

Channels are often used in conjunction with goroutines to implement concurrent algorithms. For example, you might have one goroutine that
generates values and sends them into a channel, and another goroutine that receives values from the channel and processes them. By using
channels, you can coordinate the work of these two goroutines without having to use locks or other synchronization mechanisms.

## Deadlocks

- If a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write data on that channel,
  else the program will panic and arise condition of deadlock.

## One-way or Unidirectional channels

- Golang also provides unidirectional channel where you can either only read from or only write to.

```go
readOnlyChannel := make(<-chan int)     ## a read only channel
writeOnlyChannel := make(chan<- int)     ## a write only channel
```