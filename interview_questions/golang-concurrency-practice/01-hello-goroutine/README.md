# 01 - Hello from Goroutine

## Problem
Launch a goroutine that prints `"Hello from Goroutine!"` and make sure that the main function waits for the goroutine to finish.

## Requirements
- Use the `go` keyword to create a goroutine.
- Use `sync.WaitGroup` to ensure the main function waits for the goroutine to complete.

## Example Output
```
Hello from Goroutine!
```