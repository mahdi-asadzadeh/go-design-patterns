package main

import (
	"fmt"
	"strings"
)

var allNames []string

type User struct {
	names []uint8
}

func NewUser(fullName string) *User {

	getOrAdd := func(s string) uint8 {
		for i := range allNames {
			if allNames[i] == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}

	result := User{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts {
		result.names = append(result.names, getOrAdd(p))
	}
	return &result
}

func (u *User) FullName() string {
	var parts []string
	for _, s := range u.names {
		parts = append(parts, allNames[s])
	}
	return strings.Join(parts, " ")
}

func main() {

	john2 := NewUser("John Doe")
	jane2 := NewUser("Jane Doe")
	alsoJane2 := NewUser("Jane Smith")

	fmt.Println(john2.FullName())
	fmt.Println(jane2.FullName())
	fmt.Println(alsoJane2.FullName())
}
