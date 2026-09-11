package main

import "fmt"

// filter decides whether a permutation should be included in the results.
func filter(n int) bool {
	return true
}

// permute generates all unique permutations of digits that pass filter.
func permute(digits []int, k int, results []int, seen map[int]struct{}) []int {
	if k == 1 {
		n := digitsToInt(digits)
		if filter(n) {
			if _, ok := seen[n]; !ok {
				seen[n] = struct{}{}
				results = append(results, n)
			}
		}
		return results
	}

	for i := range k {
		results = permute(digits, k-1, results, seen)
		if k%2 == 0 {
			digits[i], digits[k-1] = digits[k-1], digits[i]
		} else {
			digits[0], digits[k-1] = digits[k-1], digits[0]
		}
	}
	return results
}

// digitsToInt combines a slice of single digits into one int.
func digitsToInt(digits []int) int {
	n := 0
	for _, d := range digits {
		n = n*10 + d
	}
	return n
}

func main() {
	digits := []int{0, 2, 4, 4}
	results := make([]int, 0)
	seen := make(map[int]struct{})
	results = permute(digits, len(digits), results, seen)

	fmt.Println(results)
}
