/*
Special Pythagorean triplet
Problem 9
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
a * a + b * b = c * c

For example, 3*3 + 4*4 = 9 + 16 = 25 = 5*5.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/
package main

import "fmt"

func isPythagoreanTriplet(a int, b int, c int) bool {
	if a*a+b*b == c*c {
		return true
	}
	return false
}

func main() {
	numOne := 1
	numTwo := 1
	numThree := 1
	for i := 1; i < 1000; i++ {
		for j := i; j < 1000; j++ {
			for k := j; k < 1000; k++ {
				if i+j+k == 1000 && isPythagoreanTriplet(i, j, k) {
					numOne = i
					numTwo = j
					numThree = k
					break
				}
			}
		}
	}
	fmt.Printf("Pythogorean triplet with sum as 1000 is %d, %d, %d\n", numOne, numTwo, numThree)
	fmt.Printf("product of pythagorean triplet is %d\n", numOne*numTwo*numThree)
}
