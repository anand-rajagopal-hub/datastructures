package heap

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

	eIndex := h.exchangeIndex(position)
	if eIndex == -1 {
		return
	}
	h.exchange(position, eIndex)
	h.sink(eIndex)
}

func (h *Heap) exchange(index1, index2 int) {
	currentItem := h.items[index1]
	h.items[index1] = h.items[index2]
	h.items[index2] = currentItem
}

func (h *Heap) exchangeIndex(index int) int {
	lcIndex := index * 2
	rcIndex := (index * 2) + 1
	size := len(h.items)
	currentItem := h.items[index]
	if lcIndex >= size {
		return -1
	}

	if rcIndex >= size {
		if currentItem < h.items[lcIndex] {
			return lcIndex
		}
		return -1
	}

	if h.items[rcIndex] > h.items[lcIndex] && currentItem < h.items[rcIndex] {
		return rcIndex
	}

	if currentItem < h.items[lcIndex] {
		return lcIndex
	}
	return -1
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

func (h *Heap) insertMany(numbers ...int) {
	for _, k := range numbers {
		h.insert(k)
	}
}
