package dto

type TransactionDetailData struct {
	ID            int    `json:"id"`
	TransactionId int    `json:"transactionId" `
	CustomerName  string `json:"customerName" db:"costumer_name"`
	CustomerPhone string `json:"customerPhone" db:"costumer_phone"`
	Seat          string `json:"seat"`
}
type TransactionResponses struct {
	Id                int      `json:"id"`
	Time              string   `json:"time"`
	Location          string   `json:"location"`
	Date              string   `json:"date"`
	PriceTotal        int      `json:"priceTotal"`
	UserId            int      `json:"userId"`
	MovieTitle        string   `json:"movieTitle"`
	CinemaName        string   `json:"cinemaName"`
	PaymentMethodName string   `json:"paymentMethodName"`
	CustomerName      string   `json:"customerName" db:"costumer_name"`
	CustomerPhone     string   `json:"customerPhone" db:"costumer_phone"`
	Seat              []string `json:"seat"`
}
