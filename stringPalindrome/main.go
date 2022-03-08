package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(value string) bool {
	valueLen := len(value)
	for idx := 0; idx < valueLen; idx++ {
		if string(value[idx]) != string(value[valueLen-1-idx]) {
			return false
		}
	}
	return true
}

func getSummatedNumWithItsReversedNum(valueStr string) int {

	var revNumStr string

	// Reversing the num
	for idx := len(valueStr) - 1; idx >= 0; idx-- {
		revNumStr = revNumStr + string(valueStr[idx])
	}

	// returns Num added With Its Reversed Num
	var num, revNum int
	num, _ = strconv.Atoi(valueStr)
	revNum, _ = strconv.Atoi(revNumStr)
	return num + revNum
}

func main() {

	// Get input
	num := "195"
	itrCount := 0

	// Check If Palindrome
	for !isPalindrome(num) {
		itrCount++
		num = strconv.Itoa(getSummatedNumWithItsReversedNum(num))
	}

	fmt.Println(num, itrCount)

}
