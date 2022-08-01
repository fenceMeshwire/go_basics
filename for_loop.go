package main

import (
  "fmt"
)

func main() {
  for i := 1; i <= 20; i++ {
    if i%2 == 0 && i%3 == 0 {
      fmt.Println("Case A and Case B, with number: ", i)
      continue
    }
    if i%2 == 0 {
      fmt.Println("Case A, with number: ", i)
      continue
    }
    if i%3 == 0 {
      fmt.Println("Case B, with number: ", i)
      continue
    }
    fmt.Println("Number without case: ", i)
  }
}
