package dto

import "time"

type TransactionDetailData struct {
	ID                int       `json:"id"`
	TransactionId     int       `json:"transactionId"`
	CustomerName      string    `json:"customerName"`
	CustomerPhone     string    `json:"customerPhone"`
	Seat              string    `json:"seat"`
	Location          string    `json:"location,omitempty"`
	Time              time.Time `json:"-"`
	Date              time.Time `json:"-"`
	TimeString        string    `json:"time"`
	DateString        string    `json:"date"`
	MovieTitle        string    `json:"movieTitle"`
	CinemaName        string    `json:"cinemaName"`
	PaymentMethodName string    `json:"paymentMethodName"`
}
type TransactionResponses struct {
	Id                int    `json:"id"`
	Time              string `json:"time"`
	Location          string `json:"location"`
	Date              string `json:"date"`
	PriceTotal        int    `json:"priceTotal"`
	UserId            int    `json:"userId"`
	MovieTitle        string `json:"movieTitle"`
	CinemaName        string `json:"cinemaName"`
	PaymentMethodName string `json:"paymentMethodName"`
}
