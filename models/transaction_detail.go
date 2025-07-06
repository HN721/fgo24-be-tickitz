package models

import (
	"context"
	"fmt"
	"weeklytickits/dto"
	"weeklytickits/utils"
)

func GetTransactionDetailWithInfoByTransactionId(transactionId int) ([]dto.TransactionDetailData, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	query := `
		SELECT 
			td.id,
			td.id_transaction,
			td.costumer_name,
			td.costumer_phone,
			td.seat,
			t.time,
			t.date,
			m.title AS movie_title,
			c.name AS cinema_name,
			pm.name AS payment_method_name
		FROM transaction_detail td
		JOIN transactions t ON td.id_transaction = t.id
		JOIN movies m ON t.movie_id = m.id
		JOIN cinema c ON t.id_cinema = c.id
		JOIN payment_method pm ON t.id_payment_method = pm.id
		WHERE td.id_transaction = $1
	`

	rows, err := conn.Query(context.Background(), query, transactionId)
	if err != nil {
		return nil, fmt.Errorf("error querying transaction detail with info: %v", err)
	}
	defer conn.Conn().Close(context.Background())

	var details []dto.TransactionDetailData
	for rows.Next() {
		var detail dto.TransactionDetailData
		err := rows.Scan(
			&detail.ID,
			&detail.TransactionId,
			&detail.CustomerName,
			&detail.CustomerPhone,
			&detail.Seat,
			&detail.Time,
			&detail.Date,
			&detail.MovieTitle,
			&detail.CinemaName,
			&detail.PaymentMethodName,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}

		detail.TimeString = detail.Time.Format("15:04:05")
		detail.DateString = detail.Date.Format("2006-01-02")

		details = append(details, detail)
	}

	return details, nil
}
