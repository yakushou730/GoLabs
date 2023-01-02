package payment

import (
	"context"
	"errors"
	"fmt"

	"github.com/stripe/stripe-go/v74/charge"

	"github.com/stripe/stripe-go/v74"

	"github.com/Rhymond/go-money"
	"github.com/stripe/stripe-go/v74/client"
)

type StripeService struct {
	stripeClient *client.API
}

func NewStripeService(apikey string) (*StripeService, error) {
	if apikey == "" {
		return nil, errors.New("API key cannot be nil")
	}

	sc := &client.API{}
	sc.Init(apikey, nil)

	return &StripeService{
		stripeClient: sc,
	}, nil
}

func (s StripeService) ChargeCard(ctx context.Context, amount money.Money, cardToken string) error {
	params := &stripe.ChargeParams{
		Params:   stripe.Params{},
		Amount:   stripe.Int64(amount.Amount()),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Source: &stripe.PaymentSourceSourceParams{
			Token: stripe.String(cardToken),
		},
	}
	_, err := charge.New(params)
	if err != nil {
		return fmt.Errorf("failed to create a charge: %w", err)
	}

	return nil
}
