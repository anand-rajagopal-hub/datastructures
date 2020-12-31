package main

import (
	"fmt"

	_ "github.com/anand-rajagopal-hub/datastructures/graphs/path"
)

func createAdjList(numCourses int, prerequisites [][]int) [][]int {
	adjList := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		p := prerequisites[i]
		adjList[p[0]] = make([]int, len(p)-1)
		adjList[p[0]] = p[1:]
	}
	return adjList
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjList := createAdjList(numCourses, prerequisites)
	fmt.Println(adjList)
	return false
}

func main() {
	canFinish(2, [][]int{{1, 0}})
}
