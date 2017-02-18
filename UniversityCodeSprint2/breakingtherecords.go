package main

import "fmt"

/*
func main() {
	fmt.Println(BreakingTheRecords())
}
*/

func BreakingTheRecords() (int, int) {
	// init input data
	var N int
	fmt.Scan(&N)

	ar := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&ar[i])
	}

	//logic
	best, worst := 0, 0
	cbest, cworst := ar[0], ar[0]
	for i := 1; i < N; i++ {
		switch {
		case ar[i] > cbest:
			best++
			cbest = ar[i]
		case ar[i] < cworst:
			worst++
			cworst = ar[i]
		}
	}

	return best, worst
}
