package chapter2

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// background story
//
// “When a lead uses our app for the first time, they must pick one of three subscription plans. These
// are basic, premium, and exclusive. Depending on which they pick determines which features they get
// access to within the app. This may change over time. Once a subscription plan has been created, we
// consider that the lead has converted to a customer, and we call them a customer until they churn. At
// this point, we call them a lead again. After 6 months, we call them a lost lead and we might target
// them with a re-engagement campaign, which could include a discount code. Once a plan is created, we
// set up a recurring payment to capture funds from the customer by direct debit.”

type UserType = int
type SubscriptionType = int

const (
	unknownUserType UserType = iota
	lead
	customer
	churned
	lostLead
)

const (
	unknownSubscriptionType SubscriptionType = iota
	basic
	premium
	exclusive
)

type UserAddRequest struct {
	UserType       UserType
	Email          string
	SubType        SubscriptionType
	PaymentDetails PaymentDetails
}

type UserModifyRequest struct {
	ID             string
	UserType       UserType
	Email          string
	SubType        SubscriptionType
	PaymentDetails PaymentDetails
}

type User struct {
	ID             string
	PaymentDetails PaymentDetails
}

type PaymentDetails struct {
	stripeTokenID string
}

type UserManager interface {
	AddUser(ctx context.Context, request UserAddRequest) (User, error)
	ModifyUser(ctx context.Context, request UserModifyRequest) (User, error)
}

// amendments
type LeadRequest struct {
	email string
}

type Lead struct {
	id string
}

type LeadCreator interface {
	CreateLead(ctx context.Context, request LeadRequest) (Lead, error)
}

type Customer struct {
	leadID string
	userID string
}

func (c *Customer) UserID() string {
	return c.userID
}

func (c *Customer) SetUserID(userID string) {
	c.userID = userID
}

type LeadConvertor interface {
	Convert(ctx context.Context, subSelection SubscriptionType) (Customer, error)
}

func (l Lead) Convert(ctx context.Context, subSelection SubscriptionType) (Customer, error) {
	// TODO implement me
	panic("implement me")
}

// open host service for marketing team
type UserHandler interface {
	IsUserSubscriptionActive(ctx context.Context, userID string) bool
}

type UserActiveResponse struct {
	IsActive bool
}

func router(u UserHandler) {
	m := mux.NewRouter()
	m.HandleFunc("/user/{userID}/subscription/active", func(w http.ResponseWriter, r *http.Request) {
		// check auth, etc

		uID := mux.Vars(r)["userID"]
		if uID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		isActive := u.IsUserSubscriptionActive(r.Context(), uID)
		b, err := json.Marshal(UserActiveResponse{IsActive: isActive})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, _ = w.Write(b)
	}).Methods(http.MethodGet)
}
