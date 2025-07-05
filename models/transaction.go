package models

import (
	"context"
	"errors"
	"fmt"
	"time"
	"weeklytickits/dto"
	"weeklytickits/utils"

	"github.com/jackc/pgx/v5"
)

type Transaction struct {
	Id              int       `json:"id"`
	Time            time.Time `json:"time"`
	Date            time.Time `json:"date"`
	PriceTotal      int       `json:"priceTotal"`
	Location        string    `json:"location"`
	UserId          int       `json:"userId"`
	MovieId         int       `json:"movieId"`
	CinemaId        int       `json:"cinemaId"`
	PaymentMethodId int       `json:"paymentMethodId"`
}

type TransactionDetail struct {
	Id                int       `json:"id"`
	Time              time.Time `json:"time"`
	Date              time.Time `json:"date"`
	Location          string    `json:"location"`
	PriceTotal        int       `json:"priceTotal"`
	UserId            int       `json:"userId"`
	MovieId           int       `json:"movieId"`
	MovieTitle        string    `json:"movieTitle"`
	CinemaId          int       `json:"cinemaId"`
	CinemaName        string    `json:"cinemaName"`
	PaymentMethodId   int       `json:"paymentMethodId"`
	PaymentMethodName string    `json:"paymentMethodName"`
}
type TransactionDetailRequest struct {
	CustomerName  string `json:"customerName"`
	CustomerPhone string `json:"customerPhone"`
	Seat          string `json:"seat"`
}

func IsSeatAvailable(ctx context.Context, tx pgx.Tx, seat string, movieId, cinemaId int, date, screeningTime time.Time) (bool, error) {
	startTime := screeningTime.Add(-2 * time.Hour)
	endTime := screeningTime.Add(2 * time.Hour)

	query := `
		SELECT td.seat
		FROM transaction_detail td
		JOIN transactions t ON td.id_transaction = t.id
		WHERE td.seat = $1
		AND t.movie_id = $2
		AND t.id_cinema = $3
		AND t.date = $4
		AND t.time BETWEEN $5 AND $6
	`

	var existingSeat string
	err := tx.QueryRow(ctx, query, seat, movieId, cinemaId, date, startTime, endTime).Scan(&existingSeat)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return true, nil // seat masih kosong
		}
		return false, err
	}

	return false, nil // seat sudah dipakai
}

func CreateTransactionWithDetails(tr Transaction, details []TransactionDetailRequest) error {
	conn, err := utils.DBConnect()
	fmt.Println(tr.Location)
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	ctx := context.Background()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	success := false
	defer func() {
		if !success {
			tx.Rollback(ctx)
		}
	}()

	var transactionId int
	queryTransaction := `
	INSERT INTO transactions (time, date, price_total, user_id, movie_id, id_cinema, id_payment_method,location)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
`
	err = tx.QueryRow(ctx, queryTransaction,
		tr.Time.Format("15:04:05"),
		tr.Date.Format("2006-01-02"),
		tr.PriceTotal,
		tr.UserId,
		tr.MovieId,
		tr.CinemaId,
		tr.PaymentMethodId,
		tr.Location,
	).Scan(&transactionId)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %v", err)
	}

	queryDetail := `
		INSERT INTO transaction_detail (id_transaction, costumer_name, costumer_phone, seat)
		VALUES ($1, $2, $3, $4)
	`
	for _, detail := range details {
		available, err := IsSeatAvailable(ctx, tx, detail.Seat, tr.MovieId, tr.CinemaId, tr.Date, tr.Time)
		if err != nil {
			return fmt.Errorf("error checking seat availability: %v", err)
		}
		if !available {
			return fmt.Errorf("seat '%s' already booked for this schedule", detail.Seat)
		}

		_, err = tx.Exec(ctx, queryDetail,
			transactionId,
			detail.CustomerName,
			detail.CustomerPhone,
			detail.Seat,
		)
		if err != nil {
			return fmt.Errorf("failed to insert transaction_detail: %v", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	success = true
	return nil
}

func GetTransactionById(id int) (*dto.TransactionResponses, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	query := `
	SELECT 
	t.id, t.time, t.date, t.price_total, t.location, t.user_id, 
	t.movie_id, m.title as movie_title,
	t.id_cinema, c.name as cinema_name,
	t.id_payment_method, pm.name as payment_method_name
FROM transactions t
LEFT JOIN movies m ON t.movie_id = m.id
LEFT JOIN cinema c ON t.id_cinema = c.id
LEFT JOIN payment_method pm ON t.id_payment_method = pm.id
WHERE t.id = $1

	`

	var tr TransactionDetail
	err = conn.QueryRow(context.Background(), query, id).Scan(
		&tr.Id, &tr.Time, &tr.Date, &tr.PriceTotal, &tr.Location, &tr.UserId,
		&tr.MovieId, &tr.MovieTitle,
		&tr.CinemaId, &tr.CinemaName,
		&tr.PaymentMethodId, &tr.PaymentMethodName,
	)
	if err != nil {
		return nil, err
	}

	response := dto.TransactionResponses{
		Id:         tr.Id,
		Time:       tr.Time.Format("15:04:05"),
		Date:       tr.Date.Format("2006-01-02"),
		Location:   tr.Location,
		PriceTotal: tr.PriceTotal,

		UserId:            tr.UserId,
		MovieTitle:        tr.MovieTitle,
		CinemaName:        tr.CinemaName,
		PaymentMethodName: tr.PaymentMethodName,
	}

	return &response, nil
}
func GetTransactionsByUserId(userId int) ([]dto.TransactionResponses, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	query := `
	SELECT 
	t.id, t.time, t.date, t.price_total, t.location, t.user_id, 
	t.movie_id, m.title as movie_title,
	t.id_cinema, c.name as cinema_name,
	t.id_payment_method, pm.name as payment_method_name
FROM transactions t
LEFT JOIN movies m ON t.movie_id = m.id
LEFT JOIN cinema c ON t.id_cinema = c.id
LEFT JOIN payment_method pm ON t.id_payment_method = pm.id
		WHERE t.user_id = $1
		ORDER BY t.date DESC, t.time DESC
	`

	rows, err := conn.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responses []dto.TransactionResponses
	for rows.Next() {
		var tr TransactionDetail
		err := rows.Scan(
			&tr.Id,
			&tr.Time,
			&tr.Date,
			&tr.PriceTotal,
			&tr.Location,
			&tr.UserId,
			&tr.MovieId,
			&tr.MovieTitle,
			&tr.CinemaId,
			&tr.CinemaName,
			&tr.PaymentMethodId,
			&tr.PaymentMethodName,
		)
		if err != nil {
			return nil, err
		}

		responses = append(responses, dto.TransactionResponses{
			Id:                tr.Id,
			Time:              tr.Time.Format("15:04:05"),
			Date:              tr.Date.Format("2006-01-02"),
			Location:          tr.Location,
			PriceTotal:        tr.PriceTotal,
			UserId:            tr.UserId,
			MovieTitle:        tr.MovieTitle,
			CinemaName:        tr.CinemaName,
			PaymentMethodName: tr.PaymentMethodName,
		})
	}

	return responses, nil
}
func GetAllTransactions() ([]dto.TransactionResponses, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	query := `
	SELECT 
	t.id, t.time, t.date, t.price_total, t.location, t.user_id, 
	t.movie_id, m.title as movie_title,
	t.id_cinema, c.name as cinema_name,
	t.id_payment_method, pm.name as payment_method_name
FROM transactions t
LEFT JOIN movies m ON t.movie_id = m.id
LEFT JOIN cinema c ON t.id_cinema = c.id
LEFT JOIN payment_method pm ON t.id_payment_method = pm.id
		ORDER BY t.date DESC, t.time DESC
	`

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []dto.TransactionResponses
	for rows.Next() {
		var tr TransactionDetail
		err := rows.Scan(
			&tr.Id,
			&tr.Time,
			&tr.Date,
			&tr.PriceTotal,
			&tr.Location,
			&tr.UserId,
			&tr.MovieId,
			&tr.MovieTitle,
			&tr.CinemaId,
			&tr.CinemaName,
			&tr.PaymentMethodId,
			&tr.PaymentMethodName,
		)
		if err != nil {
			return nil, err
		}

		resp := dto.TransactionResponses{
			Id:                tr.Id,
			Time:              tr.Time.Format("15:04:05"),
			Date:              tr.Date.Format("2006-01-02"),
			Location:          tr.Location,
			PriceTotal:        tr.PriceTotal,
			UserId:            tr.UserId,
			MovieTitle:        tr.MovieTitle,
			CinemaName:        tr.CinemaName,
			PaymentMethodName: tr.PaymentMethodName,
		}

		transactions = append(transactions, resp)
	}

	return transactions, nil
}
