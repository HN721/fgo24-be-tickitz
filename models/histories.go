package models

import (
	"context"
	"weeklytickits/dto"
	"weeklytickits/utils"

	"github.com/jackc/pgx/v5"
)

func GetHistory() ([]dto.HistoryReq, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	query := `SELECT transaction_id,status,note FROM history_transaction`
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
