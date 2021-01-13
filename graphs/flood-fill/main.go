package main

import "fmt"

func dfsVisit(image [][]int, sr, sc, currentColor, newColor int) {
	image[sr][sc] = newColor
	if sr-1 >= 0 && image[sr-1][sc] == currentColor { // check above
		dfsVisit(image, sr-1, sc, currentColor, newColor)
	}
	if sr+1 <= len(image)-1 && image[sr+1][sc] == currentColor {
		dfsVisit(image, sr+1, sc, currentColor, newColor) // check below
	}
	if sc-1 >= 0 && image[sr][sc-1] == currentColor {
		dfsVisit(image, sr, sc-1, currentColor, newColor) // check left
	}
	if sc+1 <= len(image[sr])-1 && image[sr][sc+1] == currentColor {
		dfsVisit(image, sr, sc+1, currentColor, newColor)
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	currentColor := image[sr][sc]
	if currentColor == newColor {
		return image
	}
	dfsVisit(image, sr, sc, currentColor, newColor)
	return image
}

func main() {
	image := [][]int{
		{1, 1, 1},
		{1, 1, 0},
		{1, 0, 1},
	}
	fmt.Println(floodFill(image, 1, 1, 2))
}
