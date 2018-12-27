package main

// https://www.hackerrank.com/challenges/permutation-equation/problem

// Complete the permutationEquation function below.
func permutationEquation(p []int32) []int32 {
	sorted := make([]int, len(p))
	for i, v := range p {
		sorted[v-1] = i
	}

	result := make([]int32, 0, len(p))
	for i := 0; i < len(p); i++ {
		result = append(result, int32(sorted[sorted[i]])+1)
	}
	return result
}
