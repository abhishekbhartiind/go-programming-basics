package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64 = 24
	var result = math.Sqrt(num)
	var intergerResult = math.Round(result)
	fmt.Printf("Int value %.2g result", intergerResult)
	fmt.Println("---")
	fmt.Printf("Float value %.2f result", intergerResult)
}
