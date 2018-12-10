package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var inputRe *regexp.Regexp

func init() {
	inputRe = regexp.MustCompile(`position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>`)
}

func main() {
	points := readInput()

	// Iterate over points, keeping track of the total height after each simulated
	// second. Stop when the height increases.
	prevHeight := 1<<31 - 2
	height := prevHeight - 1
	seconds := 0
	for height < prevHeight {
		minY := 1<<31 - 1
		maxY := -(1<<31 - 1)
		for _, point := range points {
			point.x += point.vx
			point.y += point.vy
			if point.y < minY {
				minY = point.y
			}
			if point.y > maxY {
				maxY = point.y
			}
		}
		prevHeight = height
		height = maxY - minY
		seconds++
	}

	// Back up one iteraction to account for the last step where the height increased.
	for _, point := range points {
		point.x -= point.vx
		point.y -= point.vy
	}

	writeFile("part1.html", pointsToHtml(points))
	fmt.Println("A: <result written to part1.html>")
	fmt.Println("B:", seconds-1)
}

type point struct {
	x  int
	y  int
	vx int
	vy int
}

func readInput() []*point {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := []*point{}
	for _, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		match := inputRe.FindStringSubmatch(line)
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		vx, _ := strconv.Atoi(match[3])
		vy, _ := strconv.Atoi(match[4])
		result = append(result, &point{x, y, vx, vy})
	}

	return result
}

func pointsToHtml(points []*point) string {
	var buffer bytes.Buffer
	size := 5
	buffer.WriteString("<html><body>\n")
	for _, point := range points {
		buffer.WriteString(fmt.Sprintf("<div style=\"position:absolute; width:%dpx; height:%dpx; background:#000; left:%dpx; top: %dpx\"></div>\n", size, size, point.x*size, point.y*size))
	}
	buffer.WriteString("</body></html>")
	return buffer.String()
}

func writeFile(filename string, output string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Fprintf(file, output)
}
