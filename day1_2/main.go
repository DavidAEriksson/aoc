package main

import (
	"fmt"
	"input"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main () {
  lines := strings.Split(input.Get("1"), "\n")
  current := 0
  slice := make([]int, 0)
  for _, line := range lines {
    if line != "" {
      curr, err := strconv.Atoi(line)
      if err != nil {
        log.Fatal(err)
      }
      current += curr
    } else {
      slice = append(slice, current)
      current = 0
    }
  }
  sort.Slice(slice, func(i, j int) bool {
    return slice[i] > slice[j]
  })

  fmt.Printf("total: %v\n", slice[0] + slice[1] + slice[2])
}
