package main

// https://www.hackerrank.com/challenges/game-of-thrones/problem

// gameOfThrones function below.
func gameOfThrones(s string) string {
	m := make(map[rune]int)
	for _, s := range s {
		m[s]++
	}

	stop := false
	for _, v := range m {
		if v%2 != 0 {
			if stop {
				return "NO"
			}
			stop = true
		}
	}
	return "YES"
}
