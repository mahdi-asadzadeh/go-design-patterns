package main

import "fmt"

type IPizza interface {
	getPrice() int
}

type VaggaMania struct{}

func (p *VaggaMania) getPrice() int {
	return 15
}

type TomatoTopping struct {
	pizza IPizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

type CheeseTopping struct {
	pizza IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

func main() {
	pizza := VaggaMania{}
	fmt.Println(pizza.getPrice())

	pizzaPriseWithTomato := TomatoTopping{pizza: &pizza}
	fmt.Println(pizzaPriseWithTomato.getPrice())

	pizzaPriseWithCheese := CheeseTopping{pizza: &pizza}
	fmt.Println(pizzaPriseWithCheese.getPrice())
}
