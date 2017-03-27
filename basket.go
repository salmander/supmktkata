package supmktkata

// Costs for stuff that costs.
type Costs interface {
	Cost() float64
}

// CostsFunc adaptor for Costs.
type CostsFunc func() float64

// Cost calls f()
func (f CostsFunc) Cost() float64 {
	return f()
}

// Sandwich to eat.
type Sandwich struct{}

// Cost of a sandwich.
func (s Sandwich) Cost() float64 {
	return 2.50
}

// Crisps to snack on.
type Crisps struct{}

// Cost of some crips.
func (c Crisps) Cost() float64 {
	return 0.50
}

// Drink to guzzle on.
type Drink struct{}

// Cost of a drink.
func (d Drink) Cost() float64 {
	return 0.99
}

// Basket is a basket of items.
type Basket struct {
	Items []Costs
}

// AddItem adds an item to a basket.
func (b *Basket) AddItem(c Costs) {
	b.Items = append(b.Items, c)
}

// Total returns the total cost of basket and handles promotions
// (see README.md for promotion details).
func (b *Basket) Total() float64 {
	total := 0.00

	// Calculate meal deals
	mealDealsTotal := calculateMealDealsTotal(b)

	// Calculate total for the remaining items
	for _, item := range b.Items {
		total += item.Cost()
	}
	return total + mealDealsTotal
}

func (b *Basket) RemoveItem(c Costs) {
	for i, v := range b.Items {
		if v.Cost() == c.Cost() {
			b.Items = append(b.Items[:i], b.Items[i+1:]...)
			return
		}
	}
}

func calculateMealDealsTotal(b *Basket) float64 {
	mealDeal, crisps, sandwiches, drinks := 0, 0, 0, 0
	c := Crisps{}
	s := Sandwich{}
	d := Drink{}

	for i := 0; i < len(b.Items); i++ {
		item := b.Items[i]
		// Count Crisps in the basket
		if item.Cost() == c.Cost() {
			crisps++
		}

		// Count Sandwiches in the basket
		if item.Cost() == s.Cost() {
			sandwiches++
		}

		// Count Drinks in the basket
		if item.Cost() == d.Cost() {
			drinks++
		}

		if crisps > 0 && sandwiches > 0 && drinks > 0 {
			mealDeal++
			// Remove one of each item from the basket
			b.RemoveItem(s)
			b.RemoveItem(c)
			b.RemoveItem(d)

			// Reset all the individual counters
			i, crisps, sandwiches, drinks = -1, 0, 0, 0
		}
	}
	return float64(mealDeal * 3)
}
