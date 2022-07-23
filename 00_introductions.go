package main    // first part, package main is needed to run the main function, s. below

import (        // second part, import modules
  "fmt"
)

var (           // third part, declaration of variables
  a int = 2
  b string = "This is a string variable."
)

func main() {   // fourth part, declaration of the main function
  fmt.Println(a)
  fmt.Println(b)
}
