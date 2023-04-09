package main

import "fmt"

type Employee struct {
  Name string
  Position string
  AnnualIncome int
}


// Recommended Functional Approach
func NewEmployeeFactory(position string,
  annualIncome int) func(name string) *Employee {
  return func(name string) *Employee {
    return &Employee{name, position, annualIncome}
  }
}


func main() {
  devFactoryFunc := NewEmployeeFactory("Developer", 60000)
  mgrFactoryFunc := NewEmployeeFactory("Manager", 80000)

  //Following SOLID (HLM should not depend on LLM), it is much simpler to call a function to, in fact, do something.
  developer := devFactoryFunc("Daniel")
  fmt.Println(developer)

  manager := mgrFactoryFunc("Bruno")
  fmt.Println(manager)
}