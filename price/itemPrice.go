package price

import "fmt"

type Item struct{
	Name string
	Price float64
	Discount int
}

func calculatePrice(item Item) float64{
	tmp float64 = item.Price
	discountAmount float64 = item.Price * (item.Discount / 100)
	return tmp - discountAmount
}

func totalPrice(items []Item) float64{
	sum float64
	for i := range items {
		sum += items[i].Price
	}
}