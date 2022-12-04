package main

import (
	"fmt"
	"strings"
	"input"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main () {
  lines := strings.Split(input.Get("3"), "\n")
  total := 0

  for i := 0; i < len(lines); i += 3 {
    for j := 0; j < len(lines[i]); j++ {
      char := string(lines[i][j])
      prio := strings.Index(chars, char)
      if strings.Contains(lines[i], char) && strings.Contains(lines[i+1], char) && strings.Contains(lines[i+2], char) {
        total += (prio + 1)
        break
      }
    }
  }
  fmt.Println(total)
}
