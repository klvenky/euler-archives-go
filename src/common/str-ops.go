package common

import "fmt"

func IsStrPalindrome(str string) bool {
	byteArr := []byte(str)
	isPal := true
	length := len(byteArr)
	maxI := length - 1
	fmt.Println("str bytes ", byteArr)
	for i, j := 0, maxI; i < j; i, j = i+1, j-1 {
		isPal = byteArr[i] == byteArr[j]
		if !isPal {
			break
		}
	}
	return isPal
}
