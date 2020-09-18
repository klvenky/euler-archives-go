/*
1000-digit Fibonacci number
Problem 25

The Fibonacci sequence is defined by the recurrence relation:
    Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

    F1 = 1
    F2 = 1
    F3 = 2
    F4 = 3
    F5 = 5
    F6 = 8
    F7 = 13
    F8 = 21
    F9 = 34
    F10 = 55
    F11 = 89
    F12 = 144

The 12th term, F12, is the first term to contain three digits.
What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
*/

package main

import (
	"fmt"
	"math/big"

	"github.com/klvenky/euler-archives-go/src/common"
)

func main() {
	maxLength := "1000"
	a := new(big.Int)
	a.SetUint64(0)
	b := new(big.Int)
	b.SetUint64(1)
	index := 0
	for quit := false; quit != true; {
		index += 1
		count := common.FindDigitCount(*b)
		countStr := count.Text(10)
		tmp := new(big.Int)
		tmp.SetUint64(0)
		match := maxLength == countStr
		if match == true {
			quit = true
			break
		}

		// fmt.Println("a: ", a, "\tb: ", b, "\ttmp:", tmp, "\tlength:", countStr, "\t matched: ", match)

		// fmt.Println("tmp init done ", tmp)
		tmp.Add(tmp, b)
		// fmt.Println("tmp added b done ", tmp)
		tmp.Add(tmp, a)
		// fmt.Println("tmp added a done ", tmp)
		a = b
		b = tmp
		// fmt.Println("After manipulation a: ", a, "\tb:", b)

	}
	fmt.Printf("The first Fibanocii number of length %s is at the index is %d\n", maxLength, index)
}
