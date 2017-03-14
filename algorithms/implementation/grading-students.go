package main

// https://www.hackerrank.com/challenges/grading

import "fmt"

/*
	func main(){
		grading()
	}
*/

func grading() {
	var N int
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		var v int
		fmt.Scan(&v)
		switch {
		case v < 38:
			fmt.Println(v)
		case v%5 == 0:
			fmt.Println(v)
		case (v+1)%5 == 0:
			fmt.Println(v + 1)
		case (v+2)%5 == 0:
			fmt.Println(v + 2)
		default:
			fmt.Println(v)
		}
	}
}
