# Range And Close
- In Go, these can be used in combination with goroutines and channels to facilitate communication and synchronization between concurrent goroutines.

- The range keyword can be used to iterate over the values received from a channel until the channel is closed. Close function is basically used to close the channel.

- The range keyword simplifies the iteration over values received from a channel, while close allows you to indicate the end of data transmission and gracefully terminate goroutines waiting on the channel.
