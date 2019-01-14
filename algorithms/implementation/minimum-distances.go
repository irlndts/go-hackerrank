package main

// https://www.hackerrank.com/challenges/minimum-distances/problem

func minimumDistances(a []int32) int32 {
	result := -1
	m := make(map[int32]int)

	for i := 0; i < len(a); i++ {
		pos, ok := m[a[i]]
		if ok {
			distance := i - pos
			if result == -1 || distance < result {
				result = distance
			}
		}
		m[a[i]] = i
	}

	return int32(result)
}
