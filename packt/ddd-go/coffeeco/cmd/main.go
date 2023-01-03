package main

import (
	coffeeco "coffeeco/internal"
	"coffeeco/internal/payment"
	"coffeeco/internal/purchase"
	"coffeeco/internal/store"
	"context"
	"log"

	"github.com/Rhymond/go-money"

	"github.com/google/uuid"
)

// CoffeeCo is a national coffee shop chain. They experienced rapid growth in the last year and have opened
// 50 new stores. Each store sells coffee and coffee-related accessories, as well as store-specific drinks. Stores
// often have individual offers, but national marketing campaigns are often run, which influence the price
// of an item too.
//
// CoffeeCo recently launched a loyalty program called CoffeeBux, which allows customers to get 1 free
// drink for every 10 they purchase. It doesn't matter which store they purchase a drink at or which they
// redeem it at.

func main() {
	ctx := context.Background()
	stripeTestAPIKey := "sk_test_4eC39HqLyjWDarjtT1zdp7dc"
	cardToken := "tok_visa"
	mongoConString := "mongodb://root:example@localhost:27017"
	csvc, err := payment.NewStripeService(stripeTestAPIKey)
	if err != nil {
		log.Fatal(err)
	}

	pRepo, err := purchase.NewMongoRepo(ctx, mongoConString)
	if err != nil {
		log.Fatal(err)
	}
	if err = pRepo.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	sRepo, err := store.NewMongoRepo(ctx, mongoConString)
	if err != nil {
		log.Fatal(err)
	}
	if err = sRepo.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	sSvc := store.NewService(sRepo)

	svc := purchase.NewService(csvc, pRepo, sSvc)

	someStoreID := uuid.New()

	pur := &purchase.Purchase{
		CardToken: &cardToken,
		Store: store.Store{
			ID: someStoreID,
		},
		ProductsToPurchase: []coffeeco.Product{
			{
				ItemName:  "item1",
				BasePrice: *money.New(3300, "USD"),
			},
		},
		PaymentMeans: payment.MEANS_CARD,
	}
	if err = svc.CompletePurchase(ctx, someStoreID, pur, nil); err != nil {
		log.Fatal(err)
	}

	log.Println("purchase was successful")
}
