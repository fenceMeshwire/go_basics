package main

import (
  "fmt"
)

func main() {
  number := 15
  if number == 10 {
    fmt.Println("Lower boundary not matching: ", number) 
  } else if number > 20 {
    fmt.Println("Upper boundary not matching: ", number)
  } else {
    fmt.Println("The number lies between upper and lower boundary: ", number)
  }
}
