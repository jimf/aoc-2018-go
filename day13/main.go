package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
)

func main() {
	// Separate input into a board and agents (carts).
	board, agents := readInput()

	var firstCollision, lastRemaining agent
	numCollided := 0

	// Loop until all but 1 cart remains.
	for numCollided < len(agents)-1 {
		// Sort carts so that they appear in a "scan" order (top-left to bottom-right).
		sort.Sort(byAgent(agents))

		// Loop over each agent:
		for i, ag := range agents {
			// Skip crashed agent.
			if ag.crashed {
				continue
			}

			// Handle movement.
			nextX, nextY := nextPos(ag)
			ag.x, ag.y = nextX, nextY
			switch board[nextY][nextX] {
			case "-", "|":
				// Nothing to do.
			case "+":
				if ag.nextTurn == left {
					ag.direction = (4 + (ag.direction - 1)) % 4
				} else if ag.nextTurn == right {
					ag.direction = (ag.direction + 1) % 4
				}
				ag.nextTurn = (ag.nextTurn + 1) % 3
			case "/":
				switch ag.direction {
				case north:
					ag.direction = east
				case east:
					ag.direction = north
				case south:
					ag.direction = west
				case west:
					ag.direction = south
				}
			case "\\":
				switch ag.direction {
				case north:
					ag.direction = west
				case east:
					ag.direction = south
				case south:
					ag.direction = east
				case west:
					ag.direction = north
				}
			}

			// Detect collisions.
			for j, ag2 := range agents {
				if i != j && !ag2.crashed && ag.x == ag2.x && ag.y == ag2.y {
					numCollided += 2
					ag.crashed = true
					ag2.crashed = true
					agents[j] = ag2
					if numCollided == 2 {
						firstCollision = ag
					}
				}
			}

			agents[i] = ag
		}
	}

	for _, ag := range agents {
		if !ag.crashed {
			lastRemaining = ag
		}
	}

	fmt.Println(fmt.Sprintf("A: %d,%d", firstCollision.x, firstCollision.y))
	fmt.Println(fmt.Sprintf("B: %d,%d", lastRemaining.x, lastRemaining.y))
}

func readInput() ([][]string, []agent) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	board := [][]string{}
	agents := []agent{}
	for _, line := range strings.Split(string(input), "\n") {
		chars := strings.Split(line, "")
		if len(chars) > 0 {
			board = append(board, strings.Split(line, ""))
		}
	}
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			switch board[row][col] {
			case "^":
				agents = append(agents, agent{x: col, y: row, direction: north, nextTurn: left, crashed: false})
				board[row][col] = "|"

			case ">":
				agents = append(agents, agent{x: col, y: row, direction: east, nextTurn: left, crashed: false})
				board[row][col] = "-"

			case "v":
				agents = append(agents, agent{x: col, y: row, direction: south, nextTurn: left, crashed: false})
				board[row][col] = "|"

			case "<":
				agents = append(agents, agent{x: col, y: row, direction: west, nextTurn: left, crashed: false})
				board[row][col] = "-"
			}
		}
	}
	return board, agents
}

const (
	left     = iota
	straight = iota
	right    = iota
)

const (
	north = iota
	east  = iota
	south = iota
	west  = iota
)

type agent struct {
	x         int
	y         int
	direction int
	nextTurn  int
	crashed   bool
}

type byAgent []agent

func (agents byAgent) Len() int {
	return len(agents)
}

func (agents byAgent) Swap(i, j int) {
	agents[i], agents[j] = agents[j], agents[i]
}

func (agents byAgent) Less(i, j int) bool {
	a := agents[i]
	b := agents[j]
	return (a.y < b.y) || (a.y == b.y && a.x < b.x)
}

func nextPos(ag agent) (int, int) {
	switch ag.direction {
	case north:
		return ag.x, ag.y - 1
	case east:
		return ag.x + 1, ag.y
	case south:
		return ag.x, ag.y + 1
	case west:
		fallthrough
	default:
		return ag.x - 1, ag.y
	}
}
