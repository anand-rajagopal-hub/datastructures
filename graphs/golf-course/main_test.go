package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		forest        [][]int
		expectedSteps int
	}{
		{
			forest: [][]int{
				{1, 2, 3},
				{0, 0, 4},
				{7, 6, 5},
			},
			expectedSteps: 6,
		},
		{
			forest: [][]int{
				{1, 2, 3},
				{0, 0, 0},
				{7, 6, 5},
			},
			expectedSteps: -1,
		},
		{
			forest: [][]int{
				{54581641, 64080174, 24346381, 69107959},
				{86374198, 61363882, 68783324, 79706116},
				{668150, 92178815, 89819108, 94701471},
				{83920491, 22724204, 46281641, 47531096},
				{89078499, 18904913, 25462145, 60813308},
			},
			expectedSteps: 57,
		},
	}
	for i, tC := range testCases {
		desc := fmt.Sprintf("Test-%d", i)
		t.Run(desc, func(t *testing.T) {
			actual := cutOffTree(tC.forest)
			if actual != tC.expectedSteps {
				t.Errorf("Test case failed. Actual = %d, Expected = %d", actual, tC.expectedSteps)
			}
		})
	}
}
