/*
Longest Collatz sequence
Problem 14

The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:
13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/
package main

import (
	"fmt"
	"time"
)

func getCollatzSequenceLength(n int) int {
	next := n
	length := 0
	for i := next; i != 1; {
		if i%2 == 0 {
			i = i / 2
		} else {
			i = (3 * i) + 1
		}
		length++
	}
	return length
}
func main() {
	max := 1000000
	longestSeq := 0
	result := 0
	start := time.Now()
	for seqStartNum := 1; seqStartNum < max; seqStartNum++ {
		length := getCollatzSequenceLength(seqStartNum)
		if length > longestSeq {
			longestSeq = length
			result = seqStartNum
		}
	}
	taken := time.Since(start)
	fmt.Println("took ", taken)
	fmt.Printf("The start number that generates the longest chain is %d and it is of length : %d\n", result, longestSeq)
}
