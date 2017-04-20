package main

// https://www.hackerrank.com/challenges/camelcase
// 3 min

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	result := 1
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			result++
		}
	}
	fmt.Println(result)
}
