package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"strconv"
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
		line, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		result.PushBack(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func main() {
	input := readInput()
	freq := 0
	a := 0
	b := 0
	foundA := false
	foundB := false
	seen := make(map[int]bool)
	seen[0] = true

	for !foundB {
		for e := input.Front(); e != nil; e = e.Next() {
			freq += e.Value.(int)

			if seen[freq] && !foundB {
				b = freq
				foundB = true
			}

			seen[freq] = true
		}

		if !foundA {
			a = freq
			foundA = true
		}
	}

	fmt.Println("A:", a)
	fmt.Println("B:", b)
}
