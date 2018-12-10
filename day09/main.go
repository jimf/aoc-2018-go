package main

import (
	"container/ring"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	numPlayers, lastMarble := readInput()
	fmt.Println("A:", runMarbleGame(numPlayers, lastMarble))
	fmt.Println("B:", runMarbleGame(numPlayers, lastMarble*100))
}

func runMarbleGame(numPlayers int, lastMarble int) int {
	scores := make([]int, numPlayers)
	marbles := ring.New(1)
	marbles.Value = 0

	for marble := 1; marble <= lastMarble; marble++ {
		if marble%23 == 0 {
			player := (marble - 1) % numPlayers
			marbles = marbles.Move(-7)
			scores[player] += marble + marbles.Value.(int)
			marbles = marbles.Prev()
			marbles.Unlink(1)
			marbles = marbles.Next()
		} else {
			marbles = marbles.Next()
			current := ring.New(1)
			current.Value = marble
			marbles = marbles.Link(current)
			marbles = marbles.Prev()
		}
	}

	highScore := 0
	for _, score := range scores {
		if score > highScore {
			highScore = score
		}
	}

	return highScore
}

func readInput() (int, int) {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var (
		numPlayers int
		lastMarble int
	)
	fmt.Sscanf(string(input), "%d players; last marble is worth %d points\n", &numPlayers, &lastMarble)

	return numPlayers, lastMarble
}
