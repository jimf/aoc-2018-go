package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	points := readInput()
	fmt.Println("A:", partA(points))
	fmt.Println("B:", partB(points))
}

func partA(points []point) int {
	topLeft, bottomRight := getBounds(points)
	areas := make([]int, len(points))

	// Loop over all points inside the containing rectangle:
	for x := topLeft.x + 1; x < bottomRight.x; x++ {
		for y := topLeft.y + 1; y <= bottomRight.y; y++ {
			// If the current location is closest to a single point, inc the area for that point.
			idx, ok := closest(point{x, y}, points)
			if ok {
				areas[idx]++
			}
		}
	}

	// Now loop over the border of the rectangle. Any closest points here will be infinite, so zero them out.
	for x := topLeft.x; x <= bottomRight.x; x++ {
		for y := topLeft.y; y <= bottomRight.y; y++ {
			if x > topLeft.x && x < bottomRight.x && y > topLeft.y && y < bottomRight.y {
				continue
			}
			idx, ok := closest(point{x, y}, points)
			if ok {
				areas[idx] = 0
			}
		}
	}

	// With finite areas determined, find the max value:
	maxArea := 0
	for _, area := range areas {
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

func partB(points []point) int {
	topLeft, bottomRight := getBounds(points)
	size := 0

	for x := topLeft.x; x <= bottomRight.x; x++ {
		for y := topLeft.y; y <= bottomRight.y; y++ {
			totalDistances := 0
			for _, p := range points {
				totalDistances += distance(point{x, y}, p)
				if totalDistances >= 10000 {
					break
				}
			}
			if totalDistances < 10000 {
				size++
			}
		}
	}

	return size
}

func readInput() []point {
	input, err := ioutil.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	result := []point{}
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		parts := strings.Split(line, ", ")
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		result = append(result, point{x, y})
	}

	return result
}

// GetBounds determines the containing rectangle around a set of points and
// returns the top-left and bottom-right points of that rectangle.
func getBounds(points []point) (point, point) {
	minX, minY := 1<<31-1, 1<<31-1
	maxX, maxY := 0, 0

	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	return point{x: minX, y: minY}, point{x: maxX, y: maxY}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Distance returns the Manhattan distance between 2 points.
func distance(p1 point, p2 point) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

// Closest examines the distance between a given point and a set of points.
// If the point is closest to exactly one other point, the index of that point
// is returned. Otherwise -1 is returned.
func closest(p point, points []point) (int, bool) {
	idx := -1
	minDist := 1<<31 - 1
	numAtDist := 0

	for i, p2 := range points {
		dist := distance(p, p2)
		if dist == 0 {
			return i, true
		} else if dist < minDist {
			idx = i
			minDist = dist
			numAtDist = 1
		} else if dist == minDist {
			numAtDist++
		}
	}

	if numAtDist > 1 || idx == -1 {
		return -1, false
	}

	return idx, true
}
