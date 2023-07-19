## Goroutines Internals

### Scheduler

Similar to the scheduler that is used by the operating system, when a Go program runs, the runtime creates a scheduler that operates at the application level. This scheduler is responsible for the management and execution of goroutines.

#### OS Thread management

When a Go program starts up, the Go runtime asks the machine (virtual or physical) how many operating system threads can run in parallel.
This is based on the number of cores that are available to the program. For each thread that can be run in parallel, the runtime creates an operating system thread (`MachineCore`) and attaches that to a data structure that represents a logical processor (`Processor`) inside the program. This P and M represent the compute power or execution context for running the Go program.

Also, a `Goroutine (G)` is created to manage the execution of instructions on a selected `MachineCore`/`Processor`. Just like an `MachineCore` manages the execution of instructions on the hardware, a `G` manages the execution of instructions on the `MachineCore`.
