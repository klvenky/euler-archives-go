package main

import (
	"fmt"
	"math"
)

func main() {
	number := 1
	out := false
	if (number % 2) == 0 {
		out = false
	} else if number < 2 {
		out = false
	} else {
		max := math.Sqrt(float64(number))
		fmt.Printf("printed %f", max)
	}
	fmt.Println(out)
}
