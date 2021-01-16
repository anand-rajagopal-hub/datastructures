package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		n        []int
		k        int
		expected int
	}{
		{
			n:        []int{3, 2, 1, 5, 6, 4},
			k:        2,
			expected: 5,
		},
		{
			n:        []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:        4,
			expected: 4,
		},
		{
			n:        []int{2, 1},
			k:        2,
			expected: 1,
		},
		{
			n:        []int{7, 6, 5, 4, 3, 2, 1},
			k:        2,
			expected: 6,
		},
	}
	for i, tC := range testCases {
		desc := fmt.Sprintf("Test-%d", i+1)
		t.Run(desc, func(t *testing.T) {
			actual := findKthLargest(tC.n, tC.k)
			if actual != tC.expected {
				t.Errorf("Failed! Actual %d, Expected %d", actual, tC.expected)
			}
		})
	}
}
