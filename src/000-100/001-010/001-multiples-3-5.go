package main

import (
	"fmt"
)

/**
* Multiples of 3 and 5
* Problem 1
* If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
* Find the sum of all the multiples of 3 or 5 below 1000.
 */

func main() {
	sum := 0
	max := 1000
	i := 0
	for i < max {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
		i += 1
	}

	fmt.Printf("Sum of natural number below 100 that are multiples of 3 and 5 is %d \n", sum)
}
