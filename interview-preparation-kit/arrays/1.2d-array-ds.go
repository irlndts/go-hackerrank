package arrays

// https://www.hackerrank.com/challenges/2d-array

func sum(arr []int32) int32 {
	var result int32
	for _, v := range arr {
		result += v
	}

	return result
}

func minValue(arr [][]int32) int32 {
	var min int32
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] < min {
				min = arr[i][j]
			}
		}
	}
	if min < 0 {
		return min * 7
	}
	return 0
}

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	result := minValue(arr)
	for i := 0; i <= len(arr)-3; i++ {
		for j := 0; j <= len(arr[0])-3; j++ {
			buf := sum(arr[i][j:j+3]) + arr[i+1][j+1] + sum(arr[i+2][j:j+3])
			if buf > result {
				result = buf
			}
		}
	}
	return result
}
