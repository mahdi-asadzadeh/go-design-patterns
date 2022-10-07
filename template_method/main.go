package main

import "fmt"

type Greeter interface {
	message() string
}

type GreeterTemplate interface {
	first() string
	third() string
	greet(Greeter) string
}

type PrincesGreeterTmp struct{}

func (p *PrincesGreeterTmp) first() string {
	return "Welcome"
}
func (p *PrincesGreeterTmp) third() string {
	return "to our palace!"
}
func (p *PrincesGreeterTmp) greet(g Greeter) string {
	return fmt.Sprintf("%s, %s %s", p.first(), g.message(), p.third())
}

type PrincesGreeter struct {
	name string
}

func (p *PrincesGreeter) message() string {
	return fmt.Sprintf("your magesty princes %s", p.name)
}

func main() {
	prGreeter := &PrincesGreeter{"Diana"}
	tmpl := PrincesGreeterTmp{}

	fmt.Println(tmpl.greet(prGreeter))
}
