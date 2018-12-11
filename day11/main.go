package main

import (
	"fmt"
)

func main() {
	input := 4455

	// Part 1:
	maxTopX := 0
	maxTopY := 0
	maxPower := -(1<<31 - 1)
	for x := 1; x <= 300-2; x++ {
		for y := 1; y <= 300-2; y++ {
			power := getSquareTotalPower(x, y, 3, input)
			if power > maxPower {
				maxPower = power
				maxTopX = x
				maxTopY = y
			}
		}
	}

	// Part2: Brute forced. Takes forever, and I don't really care.
	maxTopX2 := 0
	maxTopY2 := 0
	maxSize := 0
	maxPower2 := -(1<<31 - 1)
	for size := 1; size <= 300; size++ {
		for x := 1; x <= 300-(size-1); x++ {
			for y := 1; y <= 300-(size-1); y++ {
				power := getSquareTotalPower(x, y, size, input)
				if power > maxPower2 {
					maxPower2 = power
					maxTopX2 = x
					maxTopY2 = y
					maxSize = size
				}
			}
		}
	}

	fmt.Println(fmt.Sprintf("A: %d,%d", maxTopX, maxTopY))
	fmt.Println(fmt.Sprintf("B: %d,%d,%d", maxTopX2, maxTopY2, maxSize))
}

func getCellPowerLevel(x int, y int, serialNum int) int {
	rackId := x + 10
	result := rackId*y + serialNum
	result *= rackId
	result = ((result % 1000) - result%100) / 100
	return result - 5
}

func getSquareTotalPower(topX int, topY int, size int, serialNum int) int {
	result := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			result += getCellPowerLevel(topX+i, topY+j, serialNum)
		}
	}
	return result
}
