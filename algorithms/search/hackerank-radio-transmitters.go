package main

// https://www.hackerrank.com/challenges/hackerland-radio-transmitters

import (
	"fmt"
	"sort"
)

func main() {
	var N, K int
	fmt.Scanf("%v %v", &N, &K)

	m := make(map[int]bool)
	for i := 0; i < N; i++ {
		var buf int
		fmt.Scan(&buf)
		m[buf] = false
	}

	var keys []int

	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	result := 0
	for i := keys[0]; i <= keys[len(keys)-1]; i++ {
		_, ok := m[i]
		if ok && !m[i] {

			for j := i + K; j >= i; j-- {
				_, ok := m[j]
				if ok {
					result++
					m[j] = true

					for l := j; l <= j+K; l++ {
						_, ok := m[l]
						if ok {
							m[l] = true
						}
					}

					for l := j; l >= j-K; l-- {
						_, ok := m[l]
						if ok {
							m[l] = true
						}
					}

					break
				}
			}
		}
	}
	fmt.Println(result)
}
