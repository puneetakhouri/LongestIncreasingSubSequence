package main

import "fmt"

func main() {
	arr := []int{3, 5, 7, 2, 0, 9, 1, 10}

	longestSubsequence := FindLongestIncreasingSubsequence(arr)
	fmt.Printf("The longest subsequenc length is %d\n", longestSubsequence)
}

//FindLongestIncreasingSubsequence finds the longest increasing subsequence
//from the given array
func FindLongestIncreasingSubsequence(arr []int) int {
	longestIncreasingSubsequence := make([]int, len(arr))
	max := 1
	for i := 0; i < len(arr); i++ {
		longestIncreasingSubsequence[i] = 1
		//Since each element is the longest subsequence in itself
	}

	for i := 1; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if arr[j] < arr[i] {
				if longestIncreasingSubsequence[i] < longestIncreasingSubsequence[j]+1 {
					longestIncreasingSubsequence[i] = longestIncreasingSubsequence[j] + 1
				}
			}
		}
		if max < longestIncreasingSubsequence[i] {
			max = longestIncreasingSubsequence[i]
		}
	}

	// for i := 0; i < len(arr); i++ {
	// 	if max < longestIncreasingSubsequence[i] {
	// 		max = longestIncreasingSubsequence[i]
	// 	}
	// }
	// fmt.Println()
	return max
}
