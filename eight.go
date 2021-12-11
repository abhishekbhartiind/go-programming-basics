//Function closures
/*
Go functions may be closures.
A closure is a function value that references variables from outside its body.
The function may access and assign to the referenced variables;
in this sense the function is "bound" to the variables.

For example, the adder function returns a closure.
Each closure is bound to its own sum variable.
*/

package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := 0
	b := 1
	return func() int {
		a, b = b, a+b
		return b - a
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

	fmt.Println("\nFibonacci")

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

/**
My favorite clean way to implement iterating through the Fibonacci numbers is to use first as fi - 1,
and second as fi. The Fibonacci equation states that:

fi + 1 = fi + fi - 1

Except when we write this in code, in the next round we're incrementing i. So we're effectively doing:

fnext i = fcurrent i + fcurrent i - 1

and

fnext i - 1 = fcurrent i

The way I like to implement this in code is:

first, second = second, first + second
The first = second part corresponds to updating fnext i - 1 = fcurrent i, and the second = first + second part
corresponds to updating fnext i = fcurrent i + fcurrent i - 1.

Then all we have left to do is return the old value of first,
so we'll store it in a temp variable out of the way before doing the update. In total, we get:

// fibonacci returns a function that returns
// successive fibonacci numbers from each
// successive call
func fibonacci() func() int {
    first, second := 0, 1
    return func() int {
        ret := first
        first, second = second, first+second
        return ret
    }
}

**/
