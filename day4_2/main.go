package main

import (
	"fmt"
	"input"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(input.Get("4"), "\n")
	total := 0

	for i := 0; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			break // Last line of input is not formatted correctly
		}

		pair := strings.Split(lines[i], ",")
		first := strings.Split(pair[0], "-")
		second := strings.Split(pair[1], "-")

		first_min, _ := strconv.Atoi(first[0])
		first_max, _ := strconv.Atoi(first[1])
		second_min, _ := strconv.Atoi(second[0])
		second_max, _ := strconv.Atoi(second[1])

		if overlapping(first_min, first_max, second_min, second_max) {
			total += 1
		}
	}
	fmt.Println(total)
}

func overlapping(min1 int, max1 int, min2 int, max2 int) bool {
	if max1 >= min2 && min1 <= max2 {
		return true
	} else {
		return false
	}
}
