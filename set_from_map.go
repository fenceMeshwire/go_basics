package main

import (
  "fmt"
)

func main () {
  set := map[int]bool{}
  values := []int{9, 3, 7, 7, 2, 6, 9, 1, 3, 4, 8, 1, 5, 4}
  for _, value := range(values) {
    set[value] = true
  }
  fmt.Println(set) 
}
