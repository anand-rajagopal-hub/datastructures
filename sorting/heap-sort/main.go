package main

import (
	"fmt"
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
	if position > len(h.items) {
		return
	}
	leftPos := position * 2
	rightPos := (position * 2) + 1
	if leftPos >= len(h.items) && rightPos >= len(h.items) {
		for i := position; i < len(h.items)-1; i++ {
			h.items[position] = h.items[position+1]
		}
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
	h.items[1] = -1 * (1 << 32)
	h.sink(1)

	h.items = h.items[:len(h.items)-1]
	fmt.Println(h.items)
	return max
}

func (h *Heap) insertMany(numbers ...int) {
	for _, k := range numbers {
		h.insert(k)
	}
}

func main() {
	h := NewHeap()
	// h.insert(3).insert(2).insert(5).insert(1).insert(4).insert(6)
	// fmt.Println(h.items)
	// fmt.Println(h.max())
	// fmt.Println(h.items)
	// fmt.Println(h.max())
	// fmt.Println(h.items)
	// fmt.Println(h.max())
	// fmt.Println(h.items)

	h = NewHeap()
	h.insertMany(3, 2, 3, 1, 2, 4, 5, 5, 6)
	fmt.Println(h.items)
	fmt.Println(h.max())
	fmt.Println(h.max())
	fmt.Println(h.max())
	fmt.Println(h.max())

	h = NewHeap()
	h.insertMany(3, 2, 1, 5, 6, 4)
	fmt.Println(h.items)
	fmt.Println(h.max())
	fmt.Println(h.max())
}
