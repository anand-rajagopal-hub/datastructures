package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type Heap struct {
	items []int
}

func NewHeap() *Heap {
	return &Heap{
		items: []int{0},
	}
}

func (h *Heap) insert(num int) *Heap {
	h.items = append(h.items, num)
	h.swim(len(h.items) - 1)
	return h
}

func (h *Heap) swim(position int) {
	if position == 1 {
		return
	}
	if h.items[position] > h.items[position/2] {
		t := h.items[position/2]
		h.items[position/2] = h.items[position]
		h.items[position] = t
	}
	h.swim(position / 2)
}

func (h *Heap) sink(position int) {
	if position > len(h.items)/2 {
		return
	}
	leftPos := position * 2
	rightPos := (position * 2) + 1
	if leftPos >= len(h.items) && rightPos >= len(h.items) {
		return
	}
	t := h.items[position]

	if rightPos < len(h.items) && h.items[rightPos] > h.items[leftPos] && t < h.items[rightPos] {
		h.items[position] = h.items[rightPos]
		h.items[rightPos] = t
		h.sink(rightPos)
	} else if rightPos < len(h.items) && h.items[leftPos] > h.items[rightPos] && t < h.items[leftPos] {
		h.items[position] = h.items[leftPos]
		h.items[leftPos] = t
		h.sink(leftPos)
	} else if t < h.items[leftPos] {
		h.items[position] = h.items[leftPos]
		h.items[leftPos] = t
		h.sink(leftPos)
	}
}

func (h *Heap) max() int {
	if len(h.items) == 1 {
		return -1
	}
	max := h.items[1]
	h.items[1] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.sink(1)

	// fmt.Println(h.items)
	// h.visualize(max)
	return max
}

func (h *Heap) visualize(m int) {
	var b, b1 bytes.Buffer
	b.Write([]byte("digraph heap{\n"))
	for i, v := range h.items {
		if i == 0 {
			continue
		}
		b.Write([]byte(fmt.Sprintf("%d [label=\"%d\"];\n", i, v)))
		if i > len(h.items)/2 {
			continue
		}
		if i*2 < len(h.items) {
			b1.Write([]byte(fmt.Sprintf("%d -> %d\n", i, i*2)))
		}
		if (i*2)+1 < len(h.items) {
			b1.Write([]byte(fmt.Sprintf("%d -> %d\n", i, (i*2)+1)))
		} else {
			b1.Write([]byte(fmt.Sprintf("%d -> nil\n", i)))
		}
	}
	b1.Write([]byte("}"))
	b.Write(b1.Bytes())
	err := ioutil.WriteFile(fmt.Sprintf("tree-data-%d.dot", m), b.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
}

func findKthLargest(nums []int, k int) int {
	h := NewHeap()
	for i, v := range nums {
		h.visualize(i)
		h.insert(v)
	}
	// h.visualize(0)
	for i := 0; i < k-1; i++ {
		h.max()
	}
	return h.max()
}

func main() {
	//fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	// fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}, 20))
}
