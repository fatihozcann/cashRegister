package price

import "fmt"

type Item struct {
	Name     string
	Price    float64
	Discount int
}

func CalculatePrice(item Item) float64 {
	tmp := item.Price
	if item.Discount > 0 {
		discountAmount := item.Price * float64(item.Discount) / 100
		tmp -= discountAmount
	}
	return tmp
}

func TotalPrice(items []Item) float64 {
	sum := 0.0
	for i := range items {
		sum += items[i].Price
	}
	return sum
}

// Describable interface with a Description method
type Describable interface {
	Description() string
}

// Implementing Description method for the Item type
func (item Item) Description() string {
	if item.Discount > 0 {
		return fmt.Sprintf("%s - %.2f (%.2f with %d%% discount)", item.Name, item.Price, CalculatePrice(item), item.Discount)
	}
	return fmt.Sprintf("%s - %.2f", item.Name, item.Price)
}
