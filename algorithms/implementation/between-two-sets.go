package main

import "fmt"

func main() {
	var N, M int
	fmt.Scanf("%v %v", &N, &M)

	a := make([]int, N)
	b := make([]int, M)

	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < M; i++ {
		fmt.Scan(&b[i])
	}

	result := 0
	for i := a[len(a)-1]; i <= b[0]; i++ {
		if A(i, a) && B(i, b) {
			result++
		}
	}
	fmt.Println(result)
}

func A(a int, ar []int) bool {
	for _, v := range ar {
		if a%v != 0 {
			return false
		}
	}
	return true
}

func B(b int, ar []int) bool {
	for _, v := range ar {
		if v%b != 0 {
			return false
		}

	}
	return true
}
