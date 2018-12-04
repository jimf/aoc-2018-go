package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	guardRe  *regexp.Regexp
	minuteRe *regexp.Regexp
)

func init() {
	guardRe = regexp.MustCompile(`^.+ Guard #(\d+) begins shift$`)
	minuteRe = regexp.MustCompile(`^.\d{4}-\d\d-\d\d \d\d:(\d\d).+$`)
}

// ReadInput returns the input as a sorted array of (string) events.
func readInput() []string {
	input, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(input), "\n")
	sort.Strings(lines)
	return lines[1:] // First item is empty due to a trailing empty newline after reading the file
}

// ExtractGuardId returns the guard id for begin-shift events, -1 otherwise.
func extractGuardId(event string) int {
	match := guardRe.FindStringSubmatch(event)
	if len(match) > 0 {
		id, _ := strconv.Atoi(match[1])
		return id
	}
	return -1
}

// ExtractMinute returns the minute portion of a given event.
func extractMinute(event string) int {
	match := minuteRe.FindStringSubmatch(event)
	minute, _ := strconv.Atoi(match[1])
	return minute
}

func main() {
	lines := readInput()

	sleepiestGuardByTime := -1
	sleepiestGuardByFreq := -1
	maxMin := 0
	maxMinFreq := 0
	guardSleepTime := make(map[int]int)       // guard id -> total sleep time
	guardMinutesAsleep := make(map[int][]int) // guard id -> array of sleep minute frequencies
	guardId := -1

	// Loop over all events.
	for i := 0; i < len(lines); {
		// Attempt to extract the guard id.
		id := extractGuardId(lines[i])

		// If guard id is not extracted, process sleep/wake time.
		if id == -1 {
			sleep := extractMinute(lines[i])
			wake := extractMinute(lines[i+1])
			guardSleepTime[guardId] += wake - sleep
			for min := sleep; min < wake; min++ {
				if _, ok := guardMinutesAsleep[guardId]; !ok {
					guardMinutesAsleep[guardId] = make([]int, 60)
				}
				guardMinutesAsleep[guardId][min]++
				if guardMinutesAsleep[guardId][min] > maxMinFreq {
					maxMin = min
					maxMinFreq = guardMinutesAsleep[guardId][min]
					sleepiestGuardByFreq = guardId
				}
			}
			if sleepiestGuardByTime == -1 || guardSleepTime[sleepiestGuardByTime] < guardSleepTime[guardId] {
				sleepiestGuardByTime = guardId
			}
			i += 2

			// Otherwise update current guard state.
		} else {
			guardId = id
			i++
		}
	}

	sleepiestMinute := 0
	for i := 0; i < 60; i++ {
		if guardMinutesAsleep[sleepiestGuardByTime][sleepiestMinute] < guardMinutesAsleep[sleepiestGuardByTime][i] {
			sleepiestMinute = i
		}
	}

	fmt.Println("A:", sleepiestGuardByTime*sleepiestMinute)
	fmt.Println("B:", sleepiestGuardByFreq*maxMin)
}
