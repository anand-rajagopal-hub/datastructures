package main

import (
	"fmt"
	"math"
)

type Point struct {
	x        int
	y        int
	distance float64
}

func (p *Point) String() string {
	return fmt.Sprintf("{[x, y, d] = [%d, %d, %f]}", p.x, p.y, p.distance)
}

func partition(points []*Point, start, end int) int {
	pivot := points[start].distance
	i := start + 1
	j := end
	for {
		for i <= end && points[i].distance <= pivot {
			i++
		}
		for j > start && points[j].distance > pivot {
			j--
		}
		if i >= j {
			break
		}
		points[i], points[j] = points[j], points[i]
	}
	points[start], points[j] = points[j], points[start]
	return j
}

func quickSelect(points []*Point, start, end, k int) []*Point {
	j := partition(points, start, end)
	// return points
	if j == k {
		return points[0:j]
	}
	if j > k {
		return quickSelect(points, start, j-1, k)
	}
	return quickSelect(points, j+1, end, k)
}

func kClosest(points [][]int, K int) [][]int {
	pointsArr := make([]*Point, len(points))
	for i, p := range points {
		q1 := p[0]
		q2 := p[1]
		distance := math.Sqrt(math.Pow(float64(q1), 2) + math.Pow(float64(q2), 2))
		p := &Point{
			x:        q1,
			y:        q2,
			distance: distance,
		}
		pointsArr[i] = p
	}
	// fmt.Println(pointsArr)

	closestPoints := quickSelect(pointsArr, 0, len(pointsArr)-1, K)
	result := make([][]int, len(closestPoints))
	for i, cp := range closestPoints {
		result[i] = make([]int, 2)
		result[i][0] = cp.x
		result[i][1] = cp.y
		// fmt.Println(cp.x, cp.y, cp.distance)
	}
	return result
}

func main() {
	p := [][]int{
		{3, 3},
		{5, -1},
		{-2, 4},
	}
	fmt.Println(kClosest(p, 2))
}
