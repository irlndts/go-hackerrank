package arrays

// https://www.hackerrank.com/challenges/ctci-array-left-rotation/

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {
	return append(a[d:], a[:d]...)
}
