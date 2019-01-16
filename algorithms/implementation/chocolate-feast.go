package main

// https://www.hackerrank.com/challenges/chocolate-feast/problem

// Complete the chocolateFeast function below.
func chocolateFeast(n int32, c int32, m int32) int32 {
	result := int(n / c)
	wrappers := result

	for {
		if wrappers < int(m) {
			return int32(result)
		}
		chocolates := int(wrappers / int(m))
		result += chocolates
		wrappers = chocolates + (wrappers - chocolates*int(m))
	}
}
