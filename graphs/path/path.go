package path

type Path struct {
	visited map[int]int
	paths   map[int]bool
}

// NewPath makes a new path
func NewPath(rootVertex int) *Path {
	p := &Path{
		visited: make(map[int]int),
	}

	p.visited[rootVertex] = -1
	return p
}

// Visit the given vertex
func (p *Path) Visit(vertex, from int) {
	p.visited[vertex] = from
	p.paths[from] = true
}

// Visited returns if the vertex has been visited already
func (p *Path) Visited(vertex int) (bool, int) {
	from, ok := p.visited[vertex]
	return ok, from
}

// CycleDetected returns true if the vertex is the current path
func (p *Path) CycleDetected(vertex int) bool {
	_, ok := p.paths[vertex]
	if ok {
		return true
	}
	return false
}
