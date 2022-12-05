package main

import (
	"fmt"
	"input"
	"strings"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	lines := strings.Split(input.Get("3"), "\n")
	total := 0

	for _, line := range lines {
		if len(line) == 0 {
			break // Last line of input is not formatted correctly
		}
		length := len(line)
		first := line[0:(length / 2)]
		second := line[(length / 2):length]

		for i := 0; i < len(first); i++ {
			if strings.Contains(second, string(first[i])) {
				total += strings.Index(chars, string(first[i])) + 1
				break
			}
		}
	}

	fmt.Println(total)
}
