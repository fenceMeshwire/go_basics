package main

import (
  "fmt"
  )
  
type employee struct {
  firstname   string
  lastname    string
  profession  string
  age         int
  salary     float32
}

func main() {
  worker := employee{
    "Sam", "Sample", "Programmer", 55, 1234.56,
  }
  fmt.Println()
  fmt.Println("First name:\t", worker.firstname)
  fmt.Println("Last name:\t", worker.lastname)
  fmt.Println("Profession:\t", worker.profession)
  fmt.Println("Age:\t\t", worker.age)
  fmt.Println("Salary:\t\t", worker.salary)
}
