package models

import (
	"context"
	"fmt"
	"weeklytickits/utils"
)

type Payment struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func InsertPayment(payment Payment) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	query := `INSERT INTO payment_method (name) VALUES ($1)`
	_, err = conn.Exec(context.Background(), query, payment.Name)
	return err
}

func GetAllPayments() ([]Payment, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	rows, err := conn.Query(context.Background(), `SELECT id, name FROM payment_method`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []Payment
	for rows.Next() {
		var p Payment
		if err := rows.Scan(&p.Id, &p.Name); err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}

	return payments, nil
}

func GetPaymentByID(id int) (*Payment, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	var p Payment
	err = conn.QueryRow(context.Background(),
		`SELECT id, name FROM payment_method WHERE id = $1`, id,
	).Scan(&p.Id, &p.Name)

	if err != nil {
		return nil, fmt.Errorf("Payment method with ID %d not found: %v", id, err)
	}

	return &p, nil
}

func UpdatePayment(id int, payment Payment) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	query := `UPDATE payment_method SET name = $1 WHERE id = $2`
	_, err = conn.Exec(context.Background(), query, payment.Name, id)
	return err
}

func DeletePayment(id int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	query := `DELETE FROM payment_method WHERE id = $1`
	_, err = conn.Exec(context.Background(), query, id)
	return err
}
