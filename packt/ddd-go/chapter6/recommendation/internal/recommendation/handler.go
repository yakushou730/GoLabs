package recommendation

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/Rhymond/go-money"
)

type Handler struct {
	svc Service
}

func NewHandler(svc Service) (*Handler, error) {
	if svc == (Service{}) {
		return nil, errors.New("service cannot be empty")
	}
	return &Handler{svc: svc}, nil
}

type GetRecommendationResponse struct {
	HotelName string `json:"hotelName"`
	TotalCost struct {
		Cost     int64  `json:"cost"`
		Currency string `json:"currency"`
	} `json:"totalCost"`
}

func (h Handler) GetRecommendation(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	if location == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	from := r.URL.Query().Get("from")
	if from == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	to := r.URL.Query().Get("to")
	if to == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	budget := r.URL.Query().Get("budget")
	if budget == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	const expectedFormat = "2006-01-02"

	formattedStart, err := time.Parse(expectedFormat, from)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	formattedEnd, err := time.Parse(expectedFormat, to)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b, err := strconv.ParseInt(budget, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	budgetMon := *money.New(b, "USD")

	rec, err := h.svc.Get(r.Context(), formattedStart, formattedEnd, location, budgetMon)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(GetRecommendationResponse{
		HotelName: rec.HotelName,
		TotalCost: struct {
			Cost     int64  `json:"cost"`
			Currency string `json:"currency"`
		}{
			Cost:     rec.TripPrice.Amount(),
			Currency: "USD",
		},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)

	return
}
