package main

import (
	"fmt"
	"strings"
)

const (
	vowels    = "aeiou"
	consonant = "bcdfghjklmnpqrstvwxz"
	alphabet  = "abcdefghijklmnopqurstvwxz"
)

// Melodious password
// https://www.hackerrank.com/contests/w30/challenges/melodious-password

/*
func main() {
	var N int
	fmt.Scan(&N)

	Melodious(N, "")
}
*/

func Melodious(N int, str string) {
	for i := 0; i < len(alphabet); i++ {
		if len(str) < N {
			if len(str) != 0 &&
				((strings.Contains(vowels, string(str[len(str)-1])) && strings.Contains(vowels, string(alphabet[i]))) ||
					(strings.Contains(consonant, string(str[len(str)-1])) && strings.Contains(consonant, string(alphabet[i])))) {
				// skip...
			} else {
				buf := str + string(alphabet[i])
				fmt.Println(buf)
				Melodious(N, buf)
			}
		}
	}
}
