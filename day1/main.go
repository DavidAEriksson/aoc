package main

import (
	"fmt"
	"input"
	"log"
	"strconv"
	"strings"
)

func main() {
	lines := strings.Split(input.Get("1"), "\n")
	current := 0
	largest := 0
	for _, line := range lines {
		if line != "" {
			curr, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			current += curr
		} else {
			if current >= largest {
				largest = current
			}
			current = 0
		}
	}
	fmt.Printf("%v\n", largest)
}
