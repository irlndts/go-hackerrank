package main

// Solution for https://www.hackerrank.com/challenges/dynamic-array

/*
import "fmt"

func main() {
	var N, Q int
	fmt.Scanf("%v %v", &N, &Q)

	var lastAns = 0
	var seqList [][]int
	seqList = make([][]int, N)
	for i := 0; i < N; i++ {
		seqList[i] = make([]int, 0)
	}

	for i := 0; i < Q; i++ {
		var key, x, y int
		fmt.Scanf("%v %v %v", &key, &x, &y)
		seq := (x ^ lastAns) % N
		if key == 1 {
			seqList[seq] = append(seqList[seq], y)
		} else {
			lastAns = seqList[seq][y%len(seqList[seq])]
			fmt.Println(lastAns)
		}
	}
}
*/
