package main

import (
	"fmt"
	"input"
	"strings"
)

func main() {
	total := 0
	lines := strings.Split(input.Get("2"), "\n")
	m := map[string][]int{
		"A X": {4, 3},
		"A Y": {8, 4},
		"A Z": {3, 8},
		"B X": {1, 1},
		"B Y": {5, 5},
		"B Z": {9, 9},
		"C X": {7, 2},
		"C Y": {2, 6},
		"C Z": {6, 7},
	}

	for _, line := range lines {
		if len(line) == 0 {
			break // Last line of input is not formatted correctly
		}
		total += m[line][1]
	}
	fmt.Println(total)
}
