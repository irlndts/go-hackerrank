package arrays

import (
	"sort"
)

// https://www.hackerrank.com/challenges/minimum-swaps-2/

// Arr ...
type Arr struct {
	value int32
	index int
}

// ByValue ...
type ByValue []Arr

func (a ByValue) Len() int           { return len(a) }
func (a ByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValue) Less(i, j int) bool { return a[i].value < a[j].value }

func minimumSwaps(input []int32) int32 {
	arr := make([]Arr, 0, len(input))

	// step 1. Generate an array of sorted indexes
	for i, v := range input {
		arr = append(arr, Arr{v, i})
	}

	sort.Sort(ByValue(arr))

	idx := make([]int, 0, len(input))
	for _, ar := range arr {
		idx = append(idx, ar.index)
	}

	// step 2. Sort the array by sorted indexes
	var result int32
	for i := 0; i < len(input); i++ {
		if i == idx[i] {
			continue
		}

		input[i], input[idx[i]] = input[idx[i]], input[i]
		idx[i], idx[input[idx[i]]-1] = i, idx[i]
		result++
	}
	return result
}
