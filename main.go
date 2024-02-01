package main

import (
	"cashRegister/price"
	"fmt"
)

func main() {

	var item1 price.Item
	item1.Name = "Apple"
	item1.Price = 5.0
	item1.Discount = 10

	var item2 price.Item
	item2.Name = "Penguin"
	item2.Price = 7.5
	item2.Discount = 20

	var item3 price.Item
	item3.Name = "Window"
	item3.Price = 9.0
	item3.Discount = 0

	// Item slice called items
	items := []price.Item{item1, item2, item3}

	// CalculatePrice, calculates the price and returns the price of the item
	fmt.Printf("Price of the %s is: %.2f\n", item1.Name, price.CalculatePrice(item1))
	fmt.Printf("Price of the %s is: %.2f\n", item2.Name, price.CalculatePrice(item2))
	fmt.Printf("Price of the %s is: %.2f\n", item3.Name, price.CalculatePrice(item3))

	// TotalPrice, returns the sum of the prices in an Item type slice
	total := price.TotalPrice(items)
	fmt.Printf("Total price of the slice is: %.2f\n", total)

	// Using the Describable interface
	for _, item := range items {
		fmt.Println(item.Description())
	}
}
