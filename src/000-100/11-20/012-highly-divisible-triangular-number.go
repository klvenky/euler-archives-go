/*
Highly divisible triangular number
Problem 12

The sequence of triangle numbers is generated by adding the natural numbers. So the 7th triangle number would be 1 + 2 + 3 + 4 + 5 + 6 + 7 = 28.
The first ten terms would be:
1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

Let us list the factors of the first seven triangle numbers:
1: 1
3: 1,3
6: 1,2,3,6
10: 1,2,5,10
15: 1,3,5,15
21: 1,3,7,21
28: 1,2,4,7,14,28

We can see that 28 is the first triangle number to have over five divisors.
What is the value of the first triangle number to have over five hundred divisors?
*/
package main

import (
	"fmt"
	"math"
	"time"
)

func isFactor(n int, factor int) bool {
	return (n % factor) == 0
}

func main() {
	start := time.Now()
	// maxFactors := 5
	maxFactors := 500

	prev := 0
	for index := 1; ; index++ {
		triangularNumber := prev + index
		prev = triangularNumber
		factors := 2
		squareRoot := int(math.Sqrt(float64(triangularNumber)))
		for factor := 2; factor < squareRoot; factor++ {
			if isFactor(triangularNumber, factor) {
				factors += 2
			}

		}
		if factors > maxFactors {
			taken := time.Since(start)
			fmt.Println("took ", taken)
			fmt.Printf("First Triangular Number with %d divisors is %d\nindex: %d\n", maxFactors, triangularNumber, index)
			break
		}
	}
}
