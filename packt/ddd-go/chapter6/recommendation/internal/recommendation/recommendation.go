package recommendation

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/Rhymond/go-money"
)

type Recommendation struct {
	TripStart time.Time
	TripEnd   time.Time
	HotelName string
	Location  string
	TripPrice money.Money
}

type Option struct {
	HotelName     string
	Location      string
	PricePerNight money.Money
}

type AvailabilityGetter interface {
	GetAvailability(ctx context.Context, tripStart time.Time, tripEnd time.Time, location string) ([]Option, error)
}

type Service struct {
	Availability AvailabilityGetter
}

func NewService(availability AvailabilityGetter) (*Service, error) {
	if availability == nil {
		return nil, errors.New("availability must not be nil")
	}

	return &Service{
		Availability: availability,
	}, nil
}

func (svc *Service) Get(ctx context.Context, tripStart time.Time, tripEnd time.Time, location string, budget money.Money) (*Recommendation, error) {
	switch {
	case tripStart.IsZero():
		return nil, errors.New("trip start must not be empty")
	case tripEnd.IsZero():
		return nil, errors.New("trip end must not be empty")
	case location == "":
		return nil, errors.New("location must not be empty")
	}

	opts, err := svc.Availability.GetAvailability(ctx, tripStart, tripEnd, location)
	if err != nil {
		return nil, fmt.Errorf("error getting availability: %w", err)
	}

	tripDuration := math.Round(tripEnd.Sub(tripStart).Hours() / 24)
	lowestPrice := money.NewFromFloat(999_999_999, "USD")

	var cheapestTrip *Option
	for _, opt := range opts {
		price := opt.PricePerNight.Multiply(int64(tripDuration))
		if ok, _ := price.GreaterThan(&budget); ok {
			continue
		}
		if ok, _ := price.LessThan(lowestPrice); ok {
			lowestPrice = price
			cheapestTrip = &opt
		}
	}
	if cheapestTrip == nil {
		return nil, errors.New("no trips within budget")
	}

	return &Recommendation{
		TripStart: tripStart,
		TripEnd:   tripEnd,
		HotelName: cheapestTrip.HotelName,
		Location:  cheapestTrip.Location,
		TripPrice: *lowestPrice,
	}, nil
}
