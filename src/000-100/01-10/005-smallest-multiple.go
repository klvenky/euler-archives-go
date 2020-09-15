package main

/*
Smallest multiple
Problem 5
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/
import (
	"fmt"
	"github.com/klvenky/euler-archives-go/src/common"
	"time"
)

func main() {
	// maxInt := 10
	maxInt := 20
	result := 0
	start := time.Now()
	for i := 0; ; i += maxInt {
		count := 0
		if i == 0 {
			continue
		}
		for j := 1; j <= maxInt; j++ {
			if common.IsFactor(i, j) {
				count++
			} else {
				break
			}
		}
		if count == maxInt {
			result = i
			break
		}

	}
	taken := time.Since(start)
	fmt.Printf("smallest number that is divisible by 1-%d is %d\n", maxInt, result)
	fmt.Println("took ", taken)
}
