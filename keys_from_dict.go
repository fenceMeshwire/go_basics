// print keys from a map (dictionary)

package main

import (
  "fmt"
)

func main() {
  dictionary := map[string]bool{"Peter":true, "Brian":true, "Stewie":true, "Meg":false}  
  for name := range dictionary {
    fmt.Println(name)
  }
}
