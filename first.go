package main

import "fmt"

func main() {
	var num int = 5
	var num2 = 10
	var num3 int //default value is 0
	num3 = 15

	var abc, xyz = "abc", "xyz"
	fmt.Println("Hello, Abhishek to Go Programming")
	fmt.Println(5 + 3)
	fmt.Println(num + num2 + num3)
	fmt.Println(abc, xyz)

	a := 1
	//takes care of creating a variable & assigning 1 to it
	//var a = 1
	fmt.Println(a)

	const something = "constant"
	fmt.Println(something)

	//!: Only has for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	sum := 0
	for aaa := 0; aaa < 10; aaa++ {
		sum += aaa
	}
	fmt.Println(sum)
}
