package loyalty

import (
	coffeeco "coffeeco/internal"
	"coffeeco/internal/store"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type CoffeeBux struct {
	ID                                    uuid.UUID
	store                                 store.Store
	coffeeLover                           coffeeco.CoffeeLover
	FreeDrinksAvailable                   int
	RemainingDrinkPurchasesUntilFreeDrink int
}

func (c *CoffeeBux) AddStamp() {
	if c.RemainingDrinkPurchasesUntilFreeDrink == 1 {
		c.RemainingDrinkPurchasesUntilFreeDrink = 10
		c.FreeDrinksAvailable += 1
	} else {
		c.RemainingDrinkPurchasesUntilFreeDrink--
	}
}

func (c *CoffeeBux) Pay(ctx context.Context, purchases []coffeeco.Product) error {
	lp := len(purchases)
	if lp == 0 {
		return errors.New("nothing to buy")
	}

	if c.FreeDrinksAvailable < lp {
		return fmt.Errorf("not enough coffeeBux to cover entire purchase. Have %d, need %d", lp, c.FreeDrinksAvailable)
	}

	c.FreeDrinksAvailable = c.FreeDrinksAvailable - lp

	return nil
}
