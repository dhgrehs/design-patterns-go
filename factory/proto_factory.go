package main

import "fmt"

type Employee struct {
  Name string
  Position string
  AnnualIncome int
}

const (
  Developer = iota
  Manager
)

// functional
func NewEmployee(role int) func(name string) *Employee {
  switch role {
  case Developer:
    return func(name string) *Employee {
      return &Employee{name, "Developer", 65000}
    }
  case Manager:
    return func(name string) *Employee {
      return &Employee{name, "Manager", 95000}
    }  
  }

  return func(name string) *Employee {
    return nil
  } 
}

func main() {
  mgrFactoryFunc := NewEmployee(Manager)
  mgr := mgrFactoryFunc("Sam")
  fmt.Println(mgr)
}