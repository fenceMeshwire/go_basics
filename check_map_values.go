package main

import (
  "fmt"
)

func main() {
  dictionary := map[string]int{
    "string_1": 10,
    "string_2": 20,
    }
  
  key, value_ok := dictionary["string_1"]
  fmt.Println(key, value_ok)
  
  key, value_ok = dictionary["string_2"]
  fmt.Println(key, value_ok)
  
  key, value_ok = dictionary["string_3"]
  fmt.Println(key, value_ok)
  }
