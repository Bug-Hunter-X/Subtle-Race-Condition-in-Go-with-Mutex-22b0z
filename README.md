# Subtle Race Condition in Go with Mutex

This repository demonstrates a subtle race condition in a Go program that uses goroutines and a mutex.  Even with the mutex, the program might produce an incorrect result due to a data race that is difficult to reproduce consistently.

## Bug Description

The `bug.go` file contains a program that increments a shared variable `x` using multiple goroutines. A mutex is used to protect `x` from race conditions during concurrent access. However, a subtle race condition can still occur, leading to unexpected results.

## Solution

The `bugSolution.go` file demonstrates a corrected version of the program.  The issue has been addressed by modifying the way the data is shared and accessed by concurrent goroutines to prevent the subtle race condition from occurring, ensuring the correct result every time.

## How to Run

1. Clone this repository.
2. Navigate to the repository directory.
3. Run the following commands:
   ```bash
go run bug.go
go run bugSolution.go
```
Compare the outputs of both programs to understand the effect of the bug and its solution.