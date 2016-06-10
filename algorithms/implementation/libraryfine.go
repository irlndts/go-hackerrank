package main

import "fmt"

// https://www.hackerrank.com/challenges/library-fine

func main() {
	fmt.Println(Check())
}

func Check() int {
	var d, m, y int
	var D, M, Y int
	fmt.Scanf("%v %v %v", &d, &m, &y)
	fmt.Scanf("%v %v %v", &D, &M, &Y)

	switch {
	case y < Y:
		return 0
	case y > Y:
		return 10000
	}

	switch {
	case m < M:
		return 0
	case m > M:
		return (m - M) * 500
	}

	switch {
	case d < D:
		return 0
	case d > D:
		return (d - D) * 15
	}

	return 0
}
