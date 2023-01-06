package chapter8

import _ "github.com/golang/mock/mockgen/model"

//go:generate mockgen -package mocks -destination mocks/cookies.go chapter8 CookieStockChecker,CardCharger,EmailSender
