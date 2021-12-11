//Arrays
package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//Slice
	var s []int = primes[1:4]
	fmt.Println(s)

	//Slices are like references to arrays
	//A slice does not store any data, it just describes a section of an underlying array.

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	ab := names[0:2]
	bc := names[1:3]
	fmt.Println(ab, bc)

	bc[0] = "XXX"
	fmt.Println(ab, bc)
	fmt.Println(names)
}

/**
The type [n]T is an array of n values of type T.
var a [10]int declares a variable a as an array of ten integers.
**/

//Slice
/**
An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
flexible view into the elements of an array.
In practice, slices are much more common than arrays.
The type []T is a slice with elements of type T.
A slice is formed by specifying two indices, a low and high bound, separated by a colon:
a[low : high]
**/
