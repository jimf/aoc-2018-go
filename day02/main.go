package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
)

func readInput() *list.List {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	result := list.New()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result.PushBack(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

// Leven1 compares two strings, byte by byte. If the strings have the same
// length and differ by exactly one byte, leven1 returns the index where
// they differ. Otherwise it returns -1.
func leven1(a string, b string) int {
	if len(a) != len(b) {
		return -1
	}
	differingIdx := -1
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			if differingIdx == -1 {
				differingIdx = i
			} else {
				return -1
			}
		}
	}
	return differingIdx
}

func main() {
	input := readInput()

	// Part A:
	exactlyTwo, exactlyThree := 0, 0
	for e := input.Front(); e != nil; e = e.Next() {
		chars := make(map[rune]int)
		for _, c := range e.Value.(string) {
			chars[c] += 1
		}

		foundTwo, foundThree := false, false

		for _, v := range chars {
			if v == 2 && !foundTwo {
				exactlyTwo += 1
				foundTwo = true
			} else if v == 3 && !foundThree {
				exactlyThree += 1
				foundThree = true
			}
		}
	}

	// Part B:
	sameChars := ""
	for a := input.Front(); a != nil && sameChars == ""; a = a.Next() {
		for b := a.Next(); b != nil && sameChars == ""; b = b.Next() {
			aVal, bVal := a.Value.(string), b.Value.(string)
			differingIdx := leven1(aVal, bVal)
			if differingIdx != -1 {
				sameChars = aVal[:differingIdx] + aVal[differingIdx+1:]
			}
		}
	}

	fmt.Println("A:", exactlyTwo*exactlyThree)
	fmt.Println("B:", sameChars)
}
