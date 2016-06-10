package main

//https://www.hackerrank.com/contests/projecteuler/challenges/euler001

import "fmt"

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N int
		fmt.Scan(&N)
		fmt.Println(Result(N))
	}
}

func Result(N int) int {
	// sum all for 3
	a := int(N / 3)
	if N%3 == 0 {
		a--
	}
	result := (a * (a + 1) / 2) * 3

	// sum all for 5
	b := int(N / 5)
	if N%5 == 0 {
		b--
	}
	result += (b * (b + 1) / 2) * 5

	// minus all for 15
	c := int(N / 15)
	if N%15 == 0 {
		c--
	}
	result -= (c * (c + 1) / 2) * 15

	return result
}
