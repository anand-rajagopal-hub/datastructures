package main

import (
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc            string
		grid            [][]byte
		expectedIslands int
	}{
		{
			grid: [][]byte{
				{1, 1, 1, 1, 0},
				{1, 1, 0, 1, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0},
			},
			expectedIslands: 1,
		},
		{
			grid: [][]byte{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 1, 1},
			},
			expectedIslands: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := numIslands(tC.grid)
			if tC.expectedIslands != actual {
				t.Errorf("Test case failed for input: %v", tC.grid)
			}
		})
	}
}
