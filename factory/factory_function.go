package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

func NewPerson(age int, name string) Person {
	return Person{age, name}
}

func main() {
	p := NewPerson(2, "Ali")
	fmt.Println(p)
}
