package main

import (
	"fmt"
	"strings"
  "input"
)

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main()  {
  lines := strings.Split(input.Get("5"), "\n\n")
  crates := strings.Split(lines[0], "\n")
  stack_count := crates[len(crates)-1]

	stack := map[rune]string{}

	for _, s := range crates {
		for i, r := range s {
      if strings.Contains(chars, string(r)) {
				stack[[]rune(stack_count)[i]] += string(r)
			}
		}
	}

	for _, s := range strings.Split(strings.TrimSpace(lines[1]), "\n") {
    var crate_count int
		var from, to rune
		fmt.Sscanf(s, "move %d from %c to %c", &crate_count, &from, &to)

    stack[to] = stack[from][:crate_count] + stack[to]
		stack[from] = stack[from][crate_count:]
	}

	top := ""

	for _, r := range strings.ReplaceAll(stack_count, " ", "") {
		top += stack[r][:1]
	}

	fmt.Println(top)
}
