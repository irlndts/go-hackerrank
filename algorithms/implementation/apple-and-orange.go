package main

// https://www.hackerrank.com/challenges/apple-and-orange

import "fmt"

func main() {
	a, b := AppleAndOrange()
	fmt.Printf("%d\n%d", a, b)
}

func AppleAndOrange() (apples int, oranges int) {
	var s, t, a, b, m, n int
	fmt.Scanf("%v %v", &s, &t)
	fmt.Scanf("%v %v", &a, &b)
	fmt.Scanf("%v %v", &m, &n)

	apples = Distance(a, s, t, m)
	oranges = Distance(b, s, t, n)

	return
}

func Distance(f, s, t, m int) (fruits int) {
	for i := 0; i < m; i++ {
		var fruit int
		fmt.Scan(&fruit)
		if (f+fruit) >= s && (f+fruit) <= t {
			fruits++
		}
	}
	return
}
