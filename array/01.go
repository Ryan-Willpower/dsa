package main

import (
	"fmt"
	"math"
)

func main() {
	var size int
	fmt.Println("Size of x?")
	fmt.Scan(&size)

	x := make([]float64, size)
	y := make([]float64, size)

	for i := 0; i < size; i++ {
		fmt.Printf("x[%d]:\n", i)
		fmt.Scan(&x[i])
		fmt.Printf("y[%d]:\n", i)
		fmt.Scan(&y[i])
	}

	var total float64
	for i := 0; i < size; i++ {
		total = total + math.Pow((x[i]+y[i]), 2)
	}

	fmt.Print(total)
}
