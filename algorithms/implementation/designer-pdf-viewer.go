package main

// https://www.hackerrank.com/challenges/designer-pdf-viewer

import "fmt"

/*
func main() {
	fmt.Println(DesignerPDFViewer())
}
*/

func DesignerPDFViewer() int {
	ar := make([]int, 26)

	for i := 0; i < len(ar); i++ {
		fmt.Scan(&ar[i])
	}

	var w string
	fmt.Scan(&w)

	var result int
	for _, l := range w {
		result = Max(result, ar[l-97])
	}

	return result * len(w)
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
