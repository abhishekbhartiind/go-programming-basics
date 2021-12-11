package main

import "fmt"

var xyxy = "xyxy"

func main() {
	num1 := 10
	num2 := 20
	result := add(num1, num2)
	result1, result2 := calc(num1, num2)
	fmt.Println(result)
	fmt.Println(result1, result2)
	demo()
	fmt.Println(xyxy)
}

func add(x int, y int) int {
	return x + y
}

// func calc(num1, num2 int) (int, int) {
// 	var output1 = num1 + num2
// 	var output2 = num1 - num2
// 	return output1, output2
// }

func calc(num1, num2 int) (output1, output2 int) {
	output1 = num1 + num2
	output2 = num1 - num2
	return
}
func demo() {
	var xyxy = "abcdefghigklmnopqrstuvwxyz"
	// funtional level
	fmt.Println(xyxy)
}

/*************
Every func has a scope (functional level) or it can be package level
for package level declare a var on top of the package.

Func will give a preference to local variables
over package level variables.
************/
