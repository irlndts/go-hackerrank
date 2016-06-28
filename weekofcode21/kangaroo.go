package main

//https://www.hackerrank.com/contests/w21/challenges/kangaroo

import "fmt"

func main() {
	var x1, v1, x2, v2 int
	fmt.Scanf("%v %v %v %v", &x1, &v1, &x2, &v2)
	if x2 < x1 {
		x1, v1, x2, v2 = x2, v2, x1, v1
	}
	fmt.Println(Kangaroo(x1, v1, x2, v2))
}

func Kangaroo(x1, v1, x2, v2 int) string {
	if x1 == x2 {
		return "YES"
	}
	if v2 >= v1 {
		return "NO"
	}

	for x1 <= x2 {
		if x1 == x2 {
			return "YES"
		}
		x1 += v1
		x2 += v2
	}
	return "NO"
}
