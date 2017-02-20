package main

//https://www.hackerrank.com/contests/university-codesprint-2/challenges/separate-the-numbers

import (
	"fmt"
	"strconv"
)

func main() {
	SeparateTheNumbers()
}

func SeparateTheNumbers() {
	var N int
	fmt.Scan(&N)

	var str string
	for i := 0; i < N; i++ {
		fmt.Scan(&str)
		fmt.Println(Line(str))
	}
}

func Line(str string) (string, string) {
	if len(str) == 1 {
		return "NO", ""
	}

	for j := 1; j <= len(str)/2; j++ {
		result, val := SeparateLine(j, str)
		if result {
			return "YES", strconv.Itoa(val)
		}
	}
	return "NO", ""
}

func SeparateLine(step int, str string) (bool, int) {
	begin := true
	var tmp, first int
	nines := ""
	for i := 0; i < step; i++ {
		nines = nines + "9"
	}
	for from := 0; from < len(str); from = from + step {
		to := from + step
		if len(str) < to {
			return false, first
		}

		d, _ := strconv.Atoi(str[from:to])
		if len(str[from:to]) != len(strconv.Itoa(d)) {
			return false, first
		}

		if str[from:to] == nines {
			// next decimal
			step++
			from--
		}

		if begin {
			begin = false
			first = d
			tmp = d
			continue
		}
		if d != (tmp + 1) {
			return false, first
		}
		tmp = d
	}

	return true, first
}
