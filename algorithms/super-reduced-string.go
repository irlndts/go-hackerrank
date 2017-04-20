package main

// https://www.hackerrank.com/challenges/reduced-string
// 13 min

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	N := len(s)

	for i := 0; i < N-1; {
		i++
		if s[i] == s[i-1] {
			s = s[:i-1] + s[i+1:]
			i = 0
			N = N - 2
		}
	}

	if len(s) != 0 {
		fmt.Println(s)
	} else {
		fmt.Println("Empty String")
	}
}
