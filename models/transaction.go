package models

import (
	"context"
	"fmt"
	"time"
	"weeklytickits/utils"
)

type Transaction struct {
	Id              int       `json:"id"`
	Time            time.Time `json:"time"`
	Date            time.Time `json:"date"`
	PriceTotal      int       `json:"priceTotal"`
	UserId          int       `json:"userId"`
	MovieId         int       `json:"movieId"`
	CinemaId        int       `json:"cinemaId"`
	PaymentMethodId int       `json:"paymentMethodId"`
}

type TransactionDetail struct {
	Id                int       `json:"id"`
	Time              time.Time `json:"time"`
	Date              time.Time `json:"date"`
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

func CreateTransactionWithDetails(tr Transaction, details []TransactionDetailRequest) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	success := false
	defer func() {
		if !success {
			tx.Rollback(context.Background())
		}
	}()

	var transactionId int
	queryTransaction := `
		INSERT INTO transactions (time, date, price_total, user_id, movie_id, id_cinema, id_payment_method)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id
	`
	err = tx.QueryRow(context.Background(), queryTransaction,
		tr.Time, tr.Date, tr.PriceTotal, tr.UserId, tr.MovieId, tr.CinemaId, tr.PaymentMethodId,
	).Scan(&transactionId)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %v", err)
	}

	queryDetail := `
		INSERT INTO transaction_detail (id_transaction, costumer_name, costumer_phone, seat)
		VALUES ($1, $2, $3, $4)
	`
	for _, detail := range details {
		_, err := tx.Exec(context.Background(), queryDetail,
			transactionId, detail.CustomerName, detail.CustomerPhone, detail.Seat,
		)
		if err != nil {
			return fmt.Errorf("failed to insert transaction_detail: %v", err)
		}
	}

	if err := tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	success = true
	return nil
}

func GetTransactionById(id int) (*TransactionDetail, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	query := `
		SELECT 
			t.id, t.time, t.date, t.price_total, t.user_id, 
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
		&tr.Id, &tr.Time, &tr.Date, &tr.PriceTotal, &tr.UserId,
		&tr.MovieId, &tr.MovieTitle,
		&tr.CinemaId, &tr.CinemaName,
		&tr.PaymentMethodId, &tr.PaymentMethodName)

	if err != nil {
		return nil, err
	}

	return &tr, nil
}

func GetTransactionsByUserId(userId int) ([]TransactionDetail, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	query := `
		SELECT 
			t.id, t.time, t.date, t.price_total, t.user_id, 
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

	var transactions []TransactionDetail
	for rows.Next() {
		var tr TransactionDetail
		err := rows.Scan(
			&tr.Id, &tr.Time, &tr.Date, &tr.PriceTotal, &tr.UserId,
			&tr.MovieId, &tr.MovieTitle,
			&tr.CinemaId, &tr.CinemaName,
			&tr.PaymentMethodId, &tr.PaymentMethodName)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tr)
	}

	return transactions, nil
}

func GetAllTransactions() ([]TransactionDetail, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	query := `
		SELECT 
			t.id, t.time, t.date, t.price_total, t.user_id, 
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

	var transactions []TransactionDetail
	for rows.Next() {
		var tr TransactionDetail
		err := rows.Scan(
			&tr.Id, &tr.Time, &tr.Date, &tr.PriceTotal, &tr.UserId,
			&tr.MovieId, &tr.MovieTitle,
			&tr.CinemaId, &tr.CinemaName,
			&tr.PaymentMethodId, &tr.PaymentMethodName)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tr)
	}

	return transactions, nil
}
