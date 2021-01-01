package main

import (
	"testing"
)

func TestCanFinish(t *testing.T) {
	testCases := []struct {
		desc          string
		numCourses    int
		prerequisites [][]int
		expected      bool
	}{
		{
			desc:          "test1",
			numCourses:    2,
			prerequisites: [][]int{{1, 0}, {0, 1}},
			expected:      false,
		},
		{
			desc:          "test2",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {1, 2}},
			expected:      true,
		},
		{
			desc:          "test3",
			numCourses:    4,
			prerequisites: [][]int{{3, 0}, {0, 1}},
			expected:      true,
		},
		{
			desc:          "test4",
			numCourses:    3,
			prerequisites: [][]int{{1, 0}, {0, 2}, {2, 1}},
			expected:      false,
		},
		{
			desc:          "test5",
			numCourses:    3,
			prerequisites: [][]int{{0, 1}, {0, 2}, {1, 2}},
			expected:      true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := canFinish(tC.numCourses, tC.prerequisites)
			if tC.expected != actual {
				t.Errorf("Test case failed for inputs: numCourses = %d, prerequisites = %v", tC.numCourses, tC.prerequisites)
			} else {
				t.Logf("Test case passed for inputs: numCourses = %d, prerequisites = %v", tC.numCourses, tC.prerequisites)
			}
		})
	}
}
