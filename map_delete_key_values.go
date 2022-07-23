package main

import (
  "fmt"
)

func main() {
  dictionary := map[string]int{
    "string_1": 10,
    "string_2": 20,
  }  
  delete(dictionary, "string_2")
  
  key, value_ok = dictionary["string_2"]
  fmt.Println(key, value_ok)
}
