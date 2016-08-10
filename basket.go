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
	// Implment method
	return 0.00
}
