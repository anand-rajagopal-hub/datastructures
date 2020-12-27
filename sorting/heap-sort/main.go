package main

import (
	"fmt"

	"github.com/anand-rajagopal-hub/datastructures/heap"
)

func main() {
	h := heap.NewHeap()
	// h.insert(3).insert(2).insert(5).insert(1).insert(4).insert(6)
	// fmt.Println(h.items)
	// fmt.Println(h.max())
	// fmt.Println(h.items)
	// fmt.Println(h.max())
	// fmt.Println(h.items)
	// fmt.Println(h.max())
	// fmt.Println(h.items)

	h = heap.NewHeap()
	h.insertMany(3, 2, 3, 1, 2, 4, 5, 5, 6)
	fmt.Println(h.items)
	fmt.Println(h.max())
	fmt.Println(h.max())
	fmt.Println(h.max())
	fmt.Println(h.max())

	h = heap.NewHeap()
	h.insertMany(3, 2, 1, 5, 6, 4)
	fmt.Println(h.items)
	fmt.Println(h.max())
	fmt.Println(h.max())
}
