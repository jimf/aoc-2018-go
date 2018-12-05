package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input := readInput()

	fmt.Println("A:", reactedLength(input))
	fmt.Println("B:", findShortestPolymerLength(input))
}

func readInput() string {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(input))
}

func reactedLength(polymer string) int {
	bytes := []byte(polymer)

	// Loop over each byte in input
	for i := 0; i < len(bytes)-1; i++ {
		// Compare current byte with next byte +/- 32 (ascii upper/lowercase shift)
		for i >= 0 && i < len(bytes)-2 && (bytes[i] == bytes[i+1]+32 || bytes[i] == bytes[i+1]-32) {
			// Drop current and next elements if they are the same letter, different case.
			// I think this is idiomatic? Not sure about performance.
			bytes = append(bytes[:i], bytes[i+2:]...)
			if i > 0 {
				i--
			}
		}
	}

	return len(bytes)
}

func findShortestPolymerLength(polymer string) int {
	bytes := []byte(polymer)
	result := len(polymer)

	// For each unit (letter)...
	for i := 0; i < 26; i++ {
		// Create a copy of bytes with both polarities of the unit filtered out.
		current := []byte{}
		for _, b := range bytes {
			if b != byte(i)+65 && b != byte(i)+97 {
				current = append(current, b)
			}
		}

		// React the filtered polymer and update known shortest result as needed.
		reactLen := reactedLength(string(current))
		if reactLen < result {
			result = reactLen
		}
	}

	return result
}
