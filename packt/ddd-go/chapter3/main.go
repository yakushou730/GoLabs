package chapter3

import (
	"errors"
	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
	"time"
)

// ################################# entity #################################

// Auction is an entity to represent our auction construct
type Auction struct {
	ID            int
	startingPrice money.Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}

// UUID can be used as unique identifier
type SomeEntity struct {
	id uuid.UUID
}

func NewSomeEntity() *SomeEntity {
	id := uuid.New()
	return &SomeEntity{id: id}
}

// anemic models example (do not use it to DDD)
type AnemicAuction struct {
	id            int
	startingPrice money.Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}

func (a *AnemicAuction) GetID() int {
	return a.id
}

func (a *AnemicAuction) StartingPrice() money.Money {
	return a.startingPrice
}

func (a *AnemicAuction) SetStartingPrice(startingPrice money.Money) {
	a.startingPrice = startingPrice
}

func (a *AnemicAuction) GetSellerID() int {
	return a.sellerID
}

func (a *AnemicAuction) SetSellerID(sellerID int) {
	a.sellerID = sellerID
}

func (a *AnemicAuction) GetCreatedAt() time.Time {
	return a.createdAt
}

func (a *AnemicAuction) SetCreatedAt(createdAt time.Time) {
	a.createdAt = createdAt
}

func (a *AnemicAuction) GetAuctionStart() time.Time {
	return a.auctionStart
}

func (a *AnemicAuction) SetAuctionStart(auctionStart time.Time) {
	a.auctionStart = auctionStart
}

func (a *AnemicAuction) GetAuctionEnd() time.Time {
	return a.auctionEnd
}

func (a *AnemicAuction) SetAuctionEnd(auctionEnd time.Time) {
	a.auctionEnd = auctionEnd
}

// refactor the anemic model
type AuctionRefactored struct {
	id            int
	startingPrice money.Money
	sellerID      int
	createdAt     time.Time
	auctionStart  time.Time
	auctionEnd    time.Time
}

func (a *AuctionRefactored) GetAuctionElapseDuration() time.Duration {
	return a.auctionStart.Sub(a.auctionEnd)
}

func (a *AuctionRefactored) GetAuctionEndTimeInUTC() time.Time {
	return a.auctionEnd
}

func (a *AuctionRefactored) SetAuctionEnd(auctionEnd time.Time) error {
	if err := a.validateTimeZone(auctionEnd); err != nil {
		return err
	}
	a.auctionEnd = auctionEnd

	return nil
}

func (a *AuctionRefactored) GetAuctionStartTimeInUTC() time.Time {
	return a.auctionStart
}

func (a *AuctionRefactored) SetAuctionStartTimeInUTC(auctionStart time.Time) error {
	if err := a.validateTimeZone(auctionStart); err != nil {
		return err
	}
	a.auctionStart = auctionStart

	return nil
}

func (a *AuctionRefactored) GetId() int {
	return a.id
}

func (a *AuctionRefactored) validateTimeZone(t time.Time) error {
	tz, _ := t.Zone()
	if tz != time.UTC.String() {
		return errors.New("time zone must be UTC")
	}
	return nil
}

// ################################# value object #################################

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) Point {
	return Point{
		x: x,
		y: y,
	}
}

const (
	directionUnknown = iota
	directionNorth
	directionSouth
	directionEast
	directionWest
)

func TrackPlayer() {
	currLocation := NewPoint(3, 4)
	currLocation = move(currLocation, directionNorth)
}

func move(currLocation Point, direction int) Point {
	switch direction {
	case directionNorth:
		return NewPoint(currLocation.x, currLocation.y+1)
	case directionSouth:
		return NewPoint(currLocation.x, currLocation.y-1)
	case directionEast:
		return NewPoint(currLocation.x+1, currLocation.y)
	case directionWest:
		return NewPoint(currLocation.x-1, currLocation.y)
	default:
		// do a barrel roll
	}
	return currLocation
}

// ################################# value object #################################

type WalletItem interface {
	GetBalance() (money.Money, error)
}

type Wallet struct {
	id          uuid.UUID
	ownerID     uuid.UUID
	walletItems []WalletItem
}

func (w Wallet) GetWalletBalance() (*money.Money, error) {
	var bal *money.Money
	for _, v := range w.walletItems {
		itemBal, err := v.GetBalance()
		if err != nil {
			return nil, errors.New("failed to get balance")
		}
		bal, err = bal.Add(&itemBal)
		if err != nil {
			return nil, errors.New("failed to increment balance")
		}
	}
	return bal, nil
}

// multiple customers are trying to order the same item from a website at once
type item struct {
	name string
}

type Order struct {
	items         []item
	taxAmount     money.Money
	discount      money.Money
	paymentCardID uuid.UUID
	customerID    uuid.UUID
	// marketingOptIn bool -> this attributes is not suit for order aggregate
}
