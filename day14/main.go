package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := 890691
	board := []int{3, 7}
	boardLen := 2
	elf1 := 0
	elf2 := 1
	state := s
	part1 := ""
	part2 := -1

	addRecipe := func(recipe int) {
		board = append(board, recipe)
		state = nextState(state, recipe)
		boardLen++

		// Check part 1:
		if boardLen > input && len(part1) < 10 {
			part1 = part1 + strconv.Itoa(recipe)
		}

		// Check part 2:
		if state == s890691 && part2 == -1 {
			part2 = boardLen - 6
		}
	}

	for len(part1) < 10 || part2 == -1 {
		// Append new recipes.
		next := board[elf1] + board[elf2]
		if next >= 10 {
			addRecipe(1)
		}
		addRecipe(next % 10)

		// Advance elves.
		elf1 = (elf1 + 1 + board[elf1]) % boardLen
		elf2 = (elf2 + 1 + board[elf2]) % boardLen
	}

	fmt.Println("A:", part1)
	fmt.Println("B:", part2)
}

const (
	s       = iota
	s8      = iota
	s89     = iota
	s890    = iota
	s8906   = iota
	s89069  = iota
	s890691 = iota
)

func nextState(state int, n int) int {
	if n == 8 {
		return s8
	} else if state == s8 && n == 9 {
		return s89
	} else if state == s89 && n == 0 {
		return s890
	} else if state == s890 && n == 6 {
		return s8906
	} else if state == s8906 && n == 9 {
		return s89069
	} else if state == s89069 && n == 1 {
		return s890691
	}
	return s
}
