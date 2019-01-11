package main

// https://www.hackerrank.com/challenges/repeated-string/problem

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	var result int
	for i, l := range s {
		if int64(i) == n {
			break
		}
		if string(l) == "a" {
			result++
		}
	}

	repeats := int(int(n) / len(s))
	result *= repeats

	rest := int(n) - len(s)*repeats
	for i := 0; i < rest; i++ {
		if string(s[i]) == "a" {
			result++
		}
	}

	return int64(result)
}
