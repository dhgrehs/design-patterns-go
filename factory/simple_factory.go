package main

import "fmt"

type Person struct {
  name string
  age int
}

// For simple objects like Person, you can create a very sample factory as described below:
func NewPerson(name string, age int) *Person {
  return &Person{name, age}
}

func main() {
  // Initialize directly
  p := Person{"John", 22}
  fmt.Println(p)

  // Or use a factory
  p2 := NewPerson("Jane", 21)
  p2.age = 30
  fmt.Println(p2)
}
