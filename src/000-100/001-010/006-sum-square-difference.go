package main

import "fmt"

/* Sum square difference
Problem 6
The sum of the squares of the first ten natural numbers is 385.
The square of the sum of the first ten natural numbers is 3025.
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is
3025-385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum. */

func main() {
	max := 100
	sum := 0
	squareSum := 0
	for i := 1; i <= max; i++ {
		sum += i
		squareSum += (i * i)
	}
	fmt.Printf("Difference between SquareSum & Sum is: %d\n", (sum*sum)-squareSum)
}
