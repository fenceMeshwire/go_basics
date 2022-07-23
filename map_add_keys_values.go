package main

import (
  "fmt"
)

func main() {
  dictionary := map[string]int{
    "string_1": 10,
    "string_2": 20,
  }  
  
  // Adding a key, value pair like this:
  dictionary["string_3"] = 30
  
  // Check if it worked:
  if key, value_ok := dictionary["string_3"]; value_ok {
    fmt.Println(key, value_ok)
  }
  
}
