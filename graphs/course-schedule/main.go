package main

type Path struct {
	Visited map[int]int
	Paths   map[int]int
}

// NewPath makes a new path
func NewPath() *Path {
	p := &Path{
		Visited: make(map[int]int),
	}

	return p
}

// Visit the given vertex
func (p *Path) Visit(vertex, from int) {
	p.Visited[vertex] = from
}

// VisitedAlready returns if the vertex has been visited already
func (p *Path) VisitedAlready(vertex int) (bool, int) {
	from, ok := p.Visited[vertex]
	return ok, from
}

func createAdjList(numCourses int, prerequisites [][]int) [][]int {
	adjList := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		p := prerequisites[i]
		if adjList[p[0]] != nil {
			adjList[p[0]] = append(adjList[p[0]], p[1])
		} else {
			adjList[p[0]] = make([]int, len(p)-1)
			adjList[p[0]] = p[1:]
		}
	}
	return adjList
}

func dfs(p *Path, adjList [][]int, from int, currentPath map[int]bool) bool {

	adjacentVertices := adjList[from]
	currentPath[from] = true
	for i := 0; i < len(adjacentVertices); i++ {
		vertex := adjacentVertices[i]
		if currentPath[vertex] {
			return false
		}
		ok, _ := p.VisitedAlready(vertex)
		if !ok {
			p.Visit(vertex, from)

			if !dfs(p, adjList, vertex, currentPath) {
				return false
			}

		}
	}
	delete(currentPath, from)
	return true
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjList := createAdjList(numCourses, prerequisites)
	p := NewPath()
	for i := 0; i < len(adjList); i++ {
		ok, _ := p.VisitedAlready(i)
		currentPath := make(map[int]bool)
		if !ok {
			if i == 0 {
				p.Visit(i, -1)
			} else {
				p.Visit(i, -1*i)
			}
			currentPath[i] = true
			if !dfs(p, adjList, i, currentPath) {
				return false
			}
		}

	}
	return true
}

func main() {

}
