# go-dining-philosophers-problem

This is a solution to the dining philosophers problem in Go. <br>
The dining philosophers problem is a classic synchronization problem that is used to illustrate the challenges of avoiding deadlock and starvation in concurrent systems.



## Problem Statement

* Five silent philosophers sit at a table around a bowl of spaghetti. 
* A fork is placed between each pair of adjacent philosophers. 
* Each philosopher must alternately think and eat. However, a philosopher can only eat spaghetti when they have both left and right forks. 
* Each fork can be held by only one philosopher and so a philosopher can use the fork only if it is not being used by another philosopher. 
* After they finish eating, they need to put down both forks so that the forks can be used by others. 
* A philosopher can only take the fork on their right or the one on their left as they become available, and they cannot start eating before getting both forks.


## Running the code

To run the code, simply clone the repository and run the following command:

```bash
go run cmd/main.go
```