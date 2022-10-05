package main


type Person struct {
	StreetAddress, PostCode, City, CompanyName, Position string
	AnnualIncome int
}

type PersonBuilder struct {
	person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{person: &Person{}}
} 

type PersonAddressBuilder struct{
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder 
}