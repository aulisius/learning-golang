### Buffered Channels

You can create a buffered channel by passing a second argument to the `make()` function, which specifies the maximum number of elements that the channel can hold. For example, to create a buffered channel that holds up to 10 integers, you would use the following code:

```go
ch := make(chan int, 10)
```

### Overflowing items

If you put too many items in a channel that is of fixed size, the channel will block until there is room for the items to be added. This can cause your program to hang or deadlock if there is no other code that can read from the channel and make room for new items.

For example, if you have a channel of size 2 and you try to send 3 items to it, the third send operation will block until a receiver reads a value from the channel and makes room for the third item.

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2
ch <- 3 // blocks until a receiver reads from the channel
```

To avoid this, you can use a select statement with a default case to handle the case where a send operation would block:

```go
ch := make(chan int, 2)
ch <- 1
ch <- 2

select {
case ch <- 3:
    fmt.Println("sent value to channel")
default:
    fmt.Println("channel is full, cannot send value")
}
```

In this example, the select statement with a default case handles the case where the channel is full and a send operation would block. If the channel is not full, the value is sent to the channel as usual.
