package main

import "fmt"

type Element struct {
	value int
	row   int
	col   int
}

type Heap struct {
	capacity  int
	elements  []*Element
	itemCount int
}

func NewHeap(k int) *Heap {
	h := &Heap{
		capacity:  k + 1,
		elements:  make([]*Element, k+1),
		itemCount: 0,
	}
	h.elements[h.itemCount] = nil
	h.itemCount++
	return h
}

func (h *Heap) add(e *Element) {
	// fmt.Println("adding", e)
	if len(h.elements) == h.itemCount {
		if e.value < h.elements[1].value {
			return
		}
		h.remove()
	}
	h.elements[h.itemCount] = e
	h.swim(h.itemCount)
	h.itemCount++
}

func (h *Heap) remove() *Element {
	if len(h.elements) == 1 {
		return nil
	}
	top := h.elements[1]
	h.elements[1] = h.elements[h.itemCount-1]
	h.elements[h.itemCount-1] = nil
	h.itemCount--
	h.sink(1)
	return top
}

func (h *Heap) sink(index int) {
	if index > h.itemCount/2 {
		return
	}
	exchangeIndex := h.exchangeIndex(index)
	if exchangeIndex == -1 {
		return
	}
	h.elements[index], h.elements[exchangeIndex] = h.elements[exchangeIndex], h.elements[index]
	h.sink(exchangeIndex)
}

func (h *Heap) exchangeIndex(index int) int {
	lcIndex := index << 1
	rcIndex := (index << 1) + 1
	size := h.itemCount
	currentItem := h.elements[index].value

	if lcIndex < size && rcIndex < size {
		if h.elements[lcIndex].value < h.elements[rcIndex].value && currentItem > h.elements[lcIndex].value {
			return lcIndex
		}

		if currentItem > h.elements[rcIndex].value {
			return rcIndex
		}
	}

	if lcIndex < size && currentItem > h.elements[lcIndex].value {
		return lcIndex
	}

	return -1
}

func (h *Heap) swim(index int) {
	if index == 1 {
		return
	}
	parentIndex := index >> 1
	parent := h.elements[parentIndex]
	currentElement := h.elements[index]
	if parent.value > currentElement.value {
		h.elements[parentIndex], h.elements[index] = h.elements[index], h.elements[parentIndex]
	}
	h.swim(parentIndex)
}

type ListNode struct {
	Val   int
	Level int
	Row   int
	Col   int
	Next  *ListNode
}

type Queue struct {
	head *ListNode
	tail *ListNode
}

func (q *Queue) push(val, level, row, col int) {
	if q.head == nil {
		q.head = &ListNode{
			Val:   val,
			Level: level,
			Row:   row,
			Col:   col,
		}
		q.tail = q.head
	} else {
		q.tail.Next = &ListNode{
			Val:   val,
			Level: level,
			Row:   row,
			Col:   col,
		}
		q.tail = q.tail.Next
	}
}

func (q *Queue) pop() (int, int, int, int) {
	if q.head == nil {
		return -1, -1, -1, -1
	}
	val, level, row, col := q.head.Val, q.head.Level, q.head.Row, q.head.Col
	q.head = q.head.Next
	return val, level, row, col
}

func distance(forest [][]int, sr, sc, dr, dc int) int {
	q := &Queue{}
	q.push(forest[sr][sc], 0, sr, sc)

	visited := map[string]bool{}

	for {
		val, level, sr, sc := q.pop()
		if val == -1 {
			return -1
		}
		if val == forest[dr][dc] {
			return level
		}
		// west
		key := fmt.Sprintf("%d-%d", sr, sc-1)
		if !visited[key] && sc-1 >= 0 && forest[sr][sc-1] > 0 {
			q.push(forest[sr][sc-1], level+1, sr, sc-1)
			visited[key] = true
		}
		// north
		key = fmt.Sprintf("%d-%d", sr-1, sc)
		if !visited[key] && sr-1 >= 0 && forest[sr-1][sc] > 0 {
			q.push(forest[sr-1][sc], level+1, sr-1, sc)
			visited[key] = true
		}
		// south
		key = fmt.Sprintf("%d-%d", sr+1, sc)
		if !visited[key] && sr+1 < len(forest) && forest[sr+1][sc] > 0 {
			q.push(forest[sr+1][sc], level+1, sr+1, sc)
			visited[key] = true
		}
		// east
		key = fmt.Sprintf("%d-%d", sr, sc+1)
		if !visited[key] && sc+1 < len(forest[sr]) && forest[sr][sc+1] > 0 {
			q.push(forest[sr][sc+1], level+1, sr, sc+1)
			visited[key] = true
		}
		// level++
	}

	// return -1
}

func walk(forest [][]int, h *Heap) int {

	row := 0
	col := 0
	steps := 0
	for smallest := h.remove(); smallest != nil; smallest = h.remove() {
		// fmt.Println("=========>", smallest, h)
		cr := smallest.row
		cl := smallest.col
		if row == cr && col == cl {
			forest[smallest.row][smallest.col] = 1 // visited or no tree
			continue
		}
		if forest[smallest.row][smallest.col] <= 1 {
			continue
		}
		dist := distance(forest, row, col, cr, cl)
		forest[smallest.row][smallest.col] = 1 // visited or no tree
		// fmt.Println("////////", dist)
		if dist == -1 {
			return -1
		}
		row = cr
		col = cl
		steps = steps + dist
	}
	return steps
}

func orderTree(forest [][]int) *Heap {
	h := NewHeap(len(forest) * len(forest[0]))
	for row := 0; row < len(forest); row++ {
		for col := 0; col < len(forest[row]); col++ {
			if forest[row][col] > 0 {
				h.add(&Element{
					value: forest[row][col],
					row:   row,
					col:   col,
				})
			}
		}
	}
	return h
}

func cutOffTree(forest [][]int) int {
	h := orderTree(forest)
	steps := walk(forest, h)
	return steps
}

func main() {

}
