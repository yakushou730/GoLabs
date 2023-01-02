package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"golabs/packt/ddd-go/chapter2"
	"log"
	"net/http"
	"time"
)

// ################################# factory pattern #################################

type Car interface {
	BeepBeep()
}

type BMW struct {
	heatedSeatSubscriptionEnabled bool
}

func (b BMW) BeepBeep() {
	panic("implement me")
}

type Tesla struct {
	autoPilotEnabled bool
}

func (t Tesla) BeepBeep() {
	panic("implement me")
}

func BuildCar(carType string) (Car, error) {
	switch carType {
	case "bmw":
		return BMW{heatedSeatSubscriptionEnabled: true}, nil
	case "tesla":
		return Tesla{autoPilotEnabled: true}, nil
	default:
		return nil, errors.New("unknown car type")
	}
}

func main() {
	myCar, err := BuildCar("tesla")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", myCar)
	// do something with myCar
}

// factories can help ensure business invariants are enforced at the time of object creation
type Booking struct {
	id            uuid.UUID
	from          time.Time
	to            time.Time
	hairDresserID uuid.UUID
}

func CreatBooking(from, to time.Time, hairDresserID uuid.UUID) (*Booking, error) {
	closingTime, _ := time.Parse(time.Kitchen, "17:00pm")
	if from.After(closingTime) {
		return nil, errors.New("no appointments after closing time")
	}
	return &Booking{
		id:            uuid.New(),
		hairDresserID: uuid.New(),
		from:          from,
		to:            to,
	}, nil
}

// #################################  repository layer (with Postgres) #################################

type BookingRepository interface {
	SaveBooking(ctx context.Context, booking Booking) error
	DeleteBooking(ctx context.Context, booking Booking) error
}

type PostgresRepository struct {
	connPool *pgx.Conn
}

func NewPostgresRepository(ctx context.Context, dbConnString string) (*PostgresRepository, error) {
	conn, err := pgx.Connect(ctx, dbConnString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}
	defer conn.Close(ctx)

	return &PostgresRepository{connPool: conn}, nil
}

func (p *PostgresRepository) SaveBooking(ctx context.Context, booking Booking) error {
	_, err := p.connPool.Exec(
		ctx,
		"INSERT INTO bookings (id, from, to, hair_dresser_id) VALUES ($1, $2, $3, $4)",
		booking.id.String(),
		booking.from.String(),
		booking.to.String(),
		booking.hairDresserID.String(),
	)
	if err != nil {
		return fmt.Errorf("failed to SaveBooking: %w", err)
	}
	return nil
}

func (p *PostgresRepository) DeleteBooking(ctx context.Context, booking Booking) error {
	_, err := p.connPool.Exec(
		ctx,
		"DELETE FROM bookings WHERE id = $1",
		booking.id,
	)
	if err != nil {
		return fmt.Errorf("failed to DeletingBooking: %w", err)
	}
	return nil
}

// ################################# Domain Service #################################

// 先假設我們有以下 entity
type Product struct {
	ID             int
	InStock        bool
	InsomeonesCart bool
}

func (p *Product) CanBeBought() bool {
	return p.InStock && !p.InsomeonesCart
}

type ShoppingCart struct {
	ID          int
	Products    []Product
	IsFull      bool
	MaxCartSize int
}

// This is not correct, the business logic doesn't belong to shoppingCart
// This logic should be in domain service
//func (s *ShoppingCart) AddToCart(p Product) bool {
//	if s.IsFull {
//		return false
//	}
//
//	if s.MaxCartSize == len(s.Products) {
//		s.IsFull = true
//		return false
//	}
//
//	if p.CanBeBought() {
//		s.Products = append(s.Products, p)
//		return true
//	}
//
//	return false
//}

type CheckoutService struct {
	ShoppingCart *ShoppingCart
}

func NewCheckoutService(shoppingCart *ShoppingCart) *CheckoutService {
	return &CheckoutService{ShoppingCart: shoppingCart}
}

func (c CheckoutService) AddProductToBasket(p *Product) error {
	if c.ShoppingCart.IsFull {
		return errors.New("cannot add to cart, it's full")
	}
	if p.CanBeBought() {
		c.ShoppingCart.Products = append(c.ShoppingCart.Products, *p)
		return nil
	}
	if c.ShoppingCart.MaxCartSize == len(c.ShoppingCart.Products) {
		c.ShoppingCart.IsFull = true
	}
	return nil
}

// ################################# Application Service #################################

type accountKey = int

const accountCtxKey = accountKey(1)

type BookingDomainService interface {
	CreateBooking(ctx context.Context, booking Booking) error
}

type BookingAppService struct {
	bookingRepo          BookingRepository
	bookingDomainService BookingDomainService

	// add email service
	emailService EmailSender
}

func NewBookingAppService(bookingRepo BookingRepository, bookingDomainService BookingDomainService) *BookingAppService {
	return &BookingAppService{
		bookingRepo:          bookingRepo,
		bookingDomainService: bookingDomainService,
	}
}

func (b *BookingAppService) CreateBooking(ctx context.Context, booking Booking) error {
	u, ok := ctx.Value(accountCtxKey).(chapter2.Customer)
	if !ok {
		return errors.New("invalid customer")
	}
	if u.UserID() != booking.hairDresserID.String() {
		return errors.New("cannot create booking for other users")
	}
	if err := b.bookingDomainService.CreateBooking(ctx, booking); err != nil {
		return fmt.Errorf("could not create booking: %w", err)
	}
	if err := b.bookingRepo.SaveBooking(ctx, booking); err != nil {
		return fmt.Errorf("could not save booking: %w", err)
	}

	// add email function
	err := b.emailService.SendEmail(ctx, "...", "...", "...")
	if err != nil {
		// handle it
	}

	return nil
}

// ################################# Infrastructure Service #################################

type EmailSender interface {
	SendEmail(ctx context.Context, to string, title string, body string) error
}

const emailURL = "https://mandrillapp.com/api/1.0/messages/send"

type MailChimp struct {
	apiKey     string
	from       string
	httpClient http.Client
}

type MailChimpReqBody struct {
	Key     string `json:"key"`
	Message struct {
		FromEmail string `json:"from_email"`
		Subject   string `json:"subject"`
		Text      string `json:"text"`
		To        []struct {
			Email string `json:"email"`
			Type  string `json:"type"`
		} `json:"to"`
	} `json:"message"`
}

func NewMailChimp(apiKey string, from string, httpClient http.Client) *MailChimp {
	return &MailChimp{
		apiKey:     apiKey,
		from:       from,
		httpClient: httpClient,
	}
}

func (m MailChimp) SendEmail(ctx context.Context, to string, title string, body string) error {
	bod := MailChimpReqBody{
		Key: m.apiKey,
		Message: struct {
			FromEmail string `json:"from_email"`
			Subject   string `json:"subject"`
			Text      string `json:"text"`
			To        []struct {
				Email string `json:"email"`
				Type  string `json:"type"`
			} `json:"to"`
		}{
			FromEmail: m.from,
			Subject:   title,
			Text:      body,
			To: []struct {
				Email string `json:"email"`
				Type  string `json:"type"`
			}{
				{
					Email: to,
					Type:  "to",
				},
			},
		},
	}

	b, err := json.Marshal(bod)
	if err != nil {
		return fmt.Errorf("failed to marshall body: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, emailURL, bytes.NewReader(b))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	if _, err = m.httpClient.Do(req); req != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
