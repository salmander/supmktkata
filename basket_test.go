package supmktkata_test

import (
	"testing"

	"github.com/burrbd/supmktkata"
)

func TestCalculateCostOfSingleItem(t *testing.T) {
	b := &supmktkata.Basket{}
	b.AddItem(&supmktkata.Crisps{})

	assertEquals(0.50, b.Total(), t)
}

func TestCalculateCostOfTwoItems(t *testing.T) {
	b := &supmktkata.Basket{}
	b.AddItem(&supmktkata.Crisps{})
	b.AddItem(&supmktkata.Drink{})

	assertEquals(1.49, b.Total(), t)
}

func TestMealDeal(t *testing.T) {
	b := &supmktkata.Basket{}
	b.AddItem(&supmktkata.Sandwich{})
	b.AddItem(&supmktkata.Crisps{})
	b.AddItem(&supmktkata.Drink{})

	assertEquals(3.00, b.Total(), t)
}

func TestMultipleMealDeals(t *testing.T) {
	b := &supmktkata.Basket{}
	b.AddItem(&supmktkata.Sandwich{})
	b.AddItem(&supmktkata.Crisps{})
	b.AddItem(&supmktkata.Drink{})

	// add another meal deal
	b.AddItem(&supmktkata.Sandwich{})
	b.AddItem(&supmktkata.Crisps{})
	b.AddItem(&supmktkata.Drink{})

	assertEquals(6.00, b.Total(), t)
}

func TestMealDealWithExtraItems(t *testing.T) {
	b := &supmktkata.Basket{}

	// add meal deal
	b.AddItem(&supmktkata.Sandwich{})
	b.AddItem(&supmktkata.Crisps{})
	b.AddItem(&supmktkata.Drink{})

	// add extra item
	b.AddItem(&supmktkata.Drink{})

	assertEquals(3.99, b.Total(), t)
}

func TestBasketCanCalculateOddlyPricedItems(t *testing.T) {
	b := &supmktkata.Basket{}
	b.AddItem(supmktkata.CostsFunc(func() float64 {
		return 0.65
	}))

	assertEquals(0.65, b.Total(), t)
}

func assertEquals(exp, act float64, t *testing.T) {
	if exp != act {
		t.Errorf("expected %f, got %f", exp, act)
	}
}
