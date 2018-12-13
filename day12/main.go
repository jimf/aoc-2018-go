package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	state, decisions := readInput()
	offset := 0

	for iteration := 0; iteration < 20; iteration++ {
		nextState := []string{}

		for i := -2; i < len(state)+2; i++ {
			key := lookup(&state, i-2) + lookup(&state, i-1) + lookup(&state, i) + lookup(&state, i+1) + lookup(&state, i+2)
			nextState = append(nextState, decisions[key])
		}

		start := 0
		for start < len(nextState) && nextState[start] != "#" {
			start++
		}

		end := len(nextState) - 1
		for end > 0 && nextState[end] != "#" {
			end--
		}

		if start >= end {
			offset = 0
			state = []string{}
		} else {
			offset = (offset - 2) + start
			state = nextState[start : end+1]
		}
	}

	plantNumTotal := 0
	for i, val := range state {
		if val == "#" {
			plantNumTotal += i + offset
		}
	}

	fmt.Println("A:", plantNumTotal)
	fmt.Println("B: <skipped. not interesting>")
}

func readInput() ([]string, map[string]string) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	initialState := strings.Split(lines[0][15:], "")
	decisions := make(map[string]string)
	for i := 2; i < len(lines); i++ {
		parts := strings.Split(lines[i], " => ")
		if parts[1] == "#" {
			decisions[parts[0]] = parts[1]
		}
	}
	return initialState, decisions
}

func lookup(state *[]string, idx int) string {
	if idx < 0 || idx >= len(*state) {
		return "."
	}
	result := (*state)[idx]
	if result == "" {
		return "."
	}
	return result
}
