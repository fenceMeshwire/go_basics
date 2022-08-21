package main

import (
  "fmt"
)

func main() {
  for counter := 0; counter < 10; counter++ {
    switch counter {
    case 3: 
      fmt.Println("The switch/case is: 3") 
    case 5:
      fmt.Println("The switch/case is: 5")
    default:
      fmt.Println("The default value results in: ", counter)  
    }
  }
}
