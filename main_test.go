package main

import "testing"

func TestFindLongestIncreasingSubsequence(t *testing.T) {
	testData := []struct {
		arr            []int
		expectedOutput int
	}{
		{[]int{12, 4, 8, 0, 56, 12, 6, 8, 1, 99, 16}, 4},
	}

	for _, data := range testData {
		actualOutput := FindLongestIncreasingSubsequence(data.arr)
		if actualOutput != data.expectedOutput {
			t.Errorf("Error, Expected %d, got %d", data.expectedOutput, actualOutput)
		} else {
			t.Log("SUCCESS")
		}
	}
}
