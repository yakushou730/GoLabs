package transport

import (
	"net/http"
	"recommendation/recommendation/internal/recommendation"

	"github.com/gorilla/mux"
)

func NewMux(recHandler recommendation.Handler) *mux.Router {
	m := mux.NewRouter()
	m.HandleFunc("/recommendation", recHandler.GetRecommendation).Methods(http.MethodGet)

	return m
}
