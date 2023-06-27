package main

import "fmt"

type Product struct {
	title       string
	description string
	quantity    int
	UnitPrice   float32
}

func main() {
	row := Product{
		title:       "LEGO set",
		description: "4000 pieces",
		quantity:    2,
		UnitPrice:   600,
	}

	row2 := Product{
		title:       "Plushy",
		description: "plush toy",
		quantity:    3,
		UnitPrice:   5,
	}

	row3 := Product{
		title:       "thinkpad",
		description: "laptop",
		quantity:    2,
		UnitPrice:   6000,
	}

	basket := make([]Product, 0)

	basket = append(basket, row)
	basket = append(basket, row2)
	basket = append(basket, row3)

	var sum int = 0
	for i := 0; i < len(basket); i++ {
		current := basket[i]
		fmt.Println(current)
		sum += current.quantity * int(current.UnitPrice)
	}

	fmt.Println("Total", sum)
}
