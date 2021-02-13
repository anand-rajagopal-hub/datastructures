package main

import "fmt"

func rotate(matrix [][]int) {
	rotateMat(matrix)
	// maxLayer := len(matrix) - 2
	// for layer := 0; layer <= maxLayer; layer++ {
	// 	rotateMatrix(layer, matrix)
	// }
}

func rotateMat(matrix [][]int) {
	matLength := len(matrix)
	for row := 0; row <= matLength/2; row++ {
		for col := row; col < matLength-1-row; col++ {
			tr := col
			tc := matLength - row - 1
			cur := matrix[row][col]
			for {
				temp := matrix[tr][tc]

				matrix[tr][tc] = cur
				cur = temp

				ttr := tr
				tr = tc
				tc = matLength - ttr - 1
				if tr == row && tc == col {
					matrix[tr][tc] = cur
					break
				}
			}
		}
	}
}

// func rotateMatrix(layer int, matrix [][]int) {
// 	maxRotations := len(matrix) - 1 - (2 * layer)
// 	for i := 0; i < maxRotations; i++ {
// 		col := layer
// 		row := layer
// 		maxRow := len(matrix) - layer - 1
// 		maxCol := len(matrix) - layer - 1
// 		prev := matrix[row][col]
// 		swapRow := layer
// 		swapCol := layer + 1
// 		for {
// 			cur := matrix[swapRow][swapCol]
// 			matrix[swapRow][swapCol] = prev
// 			prev = cur
// 			if swapCol == maxCol && swapRow == maxRow { // bottom right corner
// 				swapCol--
// 			} else if swapCol == layer && swapRow == maxRow { // bottom left corner
// 				swapRow--
// 			} else if swapCol == maxCol {
// 				swapRow++
// 			} else if swapRow == maxRow {
// 				swapCol--
// 			} else if swapCol == layer {
// 				swapRow--
// 			} else if swapRow == layer {
// 				swapCol++
// 			}
// 			if swapRow == layer && swapCol == layer {
// 				matrix[swapRow][swapCol] = prev
// 				break
// 			}
// 		}
// 	}
// }

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(matrix)
	fmt.Println(matrix)

	matrix = [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
