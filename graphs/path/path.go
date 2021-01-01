package path

type Path struct {
	Visited map[int]int
	Paths   map[int]int
}

// NewPath makes a new path
func NewPath() *Path {
	p := &Path{
		Visited: make(map[int]int),
		Paths:   make(map[int]int),
	}

	return p
}

// Visit the given vertex
func (p *Path) Visit(vertex, from int) {
	p.Visited[vertex] = from
	p.Paths[from] = vertex
}

// VisitedAlready returns if the vertex has been visited already
func (p *Path) VisitedAlready(vertex int) (bool, int) {
	from, ok := p.Visited[vertex]
	return ok, from
}

// CycleDetected returns true if the vertex is the current path
func (p *Path) CycleDetected(vertex int) bool {
	_, ok := p.Paths[vertex]
	if ok {
		return true
	}
	return false
}
