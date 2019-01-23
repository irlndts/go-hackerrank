package main

import (
	"fmt"
	"strconv"
)

// https://www.hackerrank.com/challenges/kaprekar-numbers/problem

// Complete the kaprekarNumbers function below.
func kaprekarNumbers(p int32, q int32) {
	var result []int
	for i := int(p); i <= int(q); i++ {
		l := len(strconv.Itoa(i))
		sq := strconv.Itoa(i * i)
		left, _ := strconv.Atoi(sq[:len(sq)-l])
		right, _ := strconv.Atoi(sq[len(sq)-l:])
		if left+right == i {
			result = append(result, i)
		}
	}
	if len(result) == 0 {
		fmt.Printf("INVALID RANGE")
		return
	}

	for _, v := range result {
		fmt.Printf("%v ", v)
	}
}
