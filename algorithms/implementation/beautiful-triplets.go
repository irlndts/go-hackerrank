package main

// https://www.hackerrank.com/challenges/beautiful-triplets/problem

func beautifulTriplets(d int32, arr []int32) int32 {
	result := int32(0)
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[j]-arr[i] != d {
				continue
			}
			for k := j + 1; k < len(arr); k++ {
				if arr[k]-arr[j] == d {
					result++
				}
			}
		}
	}
	return result
}
