package main

import "fmt"

//https://www.hackerrank.com/challenges/array-left-rotation

func main() {
	LeftRotation()
}

func LeftRotation() {
	var N, d int
	fmt.Scan(&N)
	fmt.Scan(&d)

	ar := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&ar[i])
	}

	mid := d % N

	result := make([]int, N)
	result = append(ar[mid:], ar[:mid]...)
	for _, v := range result {
		fmt.Printf("%v ", v)
	}
}
