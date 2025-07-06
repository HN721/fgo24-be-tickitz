package models

import (
	"context"
	"fmt"
	"weeklytickits/dto"
	"weeklytickits/utils"

	"github.com/jackc/pgx/v5"
)

func GetHistory() ([]dto.HistoryReq, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	query := `SELECT id,transaction_id,status,note FROM history_transaction`
	rows, err := conn.Query(context.Background(), query)
	data, err := pgx.CollectRows[dto.HistoryReq](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}
	return data, nil

}
func UpdateHistory(historyId int, data dto.HistoryReq) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	query := `UPDATE history_transaction SET status =$1 ,note =$2 WHERE id = $3`
	_, err = conn.Exec(context.Background(), query, data.Status, data.Note, historyId)
	return err

}

func GetHistoryByIdUser(userId int) ([]dto.HistoryReq, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	fmt.Println(userId)
	defer conn.Conn().Close(context.Background())

	query := `SELECT ht.transaction_id, ht.status, ht.note
	FROM history_transaction ht
	JOIN transactions t ON ht.transaction_id = t.id
	WHERE t.user_id = $1`

	rows, err := conn.Query(context.Background(), query, userId)
	if err != nil {
		return nil, err
	}

	data, err := pgx.CollectRows[dto.HistoryReq](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}

	return data, nil
}
