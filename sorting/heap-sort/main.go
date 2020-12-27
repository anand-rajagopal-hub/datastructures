package main

import (
	"fmt"

	"github.com/anand-rajagopal-hub/datastructures/heap"
)

func main() {

	h = heap.NewHeap()
	h.InsertMany(3, 2, 3, 1, 2, 4, 5, 5, 6)
	h.PrintItems()
	fmt.Println(h.Max())
	fmt.Println(h.Max())
	fmt.Println(h.Max())
	fmt.Println(h.Max())

	h = heap.NewHeap()
	h.InsertMany(3, 2, 1, 5, 6, 4)
	h.PrintItems()
	fmt.Println(h.Max())
	fmt.Println(h.Max())
}
