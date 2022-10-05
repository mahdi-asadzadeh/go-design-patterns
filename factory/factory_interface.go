package main

import "fmt"

type Person interface {
	SayHello() string
}

type person struct {
	Age  int
	Name string
}

func (p *person) SayHello() string{
	return "Hello"
}

func NewPerson(age int, name string) Person {
	return &person{age, name}
}

func main(){
	p := NewPerson(2, "Ali")
	fmt.Println(p)
}