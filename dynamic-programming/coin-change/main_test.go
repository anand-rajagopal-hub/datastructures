package main

import "testing"

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		coinChange([]int{186, 419, 83, 408}, 6249)
	}
}
