package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var claimRe *regexp.Regexp

type claim struct {
	id     int
	x      int
	y      int
	width  int
	height int
}

func init() {
	claimRe = regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)
}

func parseLine(line string) claim {
	match := claimRe.FindStringSubmatch(line)
	id, _ := strconv.Atoi(match[1])
	x, _ := strconv.Atoi(match[2])
	y, _ := strconv.Atoi(match[3])
	width, _ := strconv.Atoi(match[4])
	height, _ := strconv.Atoi(match[5])
	return claim{id, x, y, width, height}
}

func readInput() *list.List {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := list.New()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result.PushBack(parseLine(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func checkOverlap(a claim, b claim) bool {
	return a.x+a.width > b.x &&
		a.x < b.x+b.width &&
		a.y+a.height > b.y &&
		a.y < b.y+b.height
}

func main() {
	claims := readInput()
	var fabric [1000][1000]int

	// Part A: Load fabric grid and record how many claims overlap.
	a := 0
	for e := claims.Front(); e != nil; e = e.Next() {
		claim := e.Value.(claim)
		for i := claim.x; i < claim.x+claim.width; i++ {
			for j := claim.y; j < claim.y+claim.height; j++ {
				fabric[i][j]++
				if fabric[i][j] == 2 {
					a++
				}
			}
		}
	}

	// Part B: Find claim that does not overlap.
	b := -1
	for claim1 := claims.Front(); claim1 != nil && b == -1; claim1 = claim1.Next() {
		claim1Val := claim1.Value.(claim)
		overlapsAny := false
		for claim2 := claims.Front(); claim2 != nil && !overlapsAny; claim2 = claim2.Next() {
			claim2Val := claim2.Value.(claim)
			if claim1Val != claim2Val && checkOverlap(claim1Val, claim2Val) {
				overlapsAny = true
			}
		}
		if !overlapsAny {
			b = claim1Val.id
		}
	}

	fmt.Println("A:", a)
	fmt.Println("B:", b)
}
