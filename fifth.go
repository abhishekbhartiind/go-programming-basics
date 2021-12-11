package main

import (
	"fmt"
)

/**
Deferred function calls are pushed onto a stack.
When a function returns, its deferred calls are executed in last-in-first-out order.
**/
func main() {
	defer fmt.Println("world")
	//defer will excute the function after the main function
	fmt.Println("hello")

	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}
