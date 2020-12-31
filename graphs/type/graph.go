package type

type Path struct {
	visited map[int]int
	paths map[int]bool
}

func NewPath(rootVertex int)  *Path {
	p := &Path{
		visited: make(map[int]int)
	}

	p[rootVertex] = -1
	return p
}

func (p *Path) visit(vertex, from int) {
	p.visited[vertex] = from
	paths[from] = true
}

func (p *Path) visited(vertex) (bool, int) {
	from, ok := p.visited[vertex]
	return ok, from
}

func (p *Path) cycleDetected(vertex int) bool  {
	_, ok := p.paths[vertex]
	if ok {
		return true
	}
	return false
}