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

	query := `
	SELECT 
		td.id,
		td.id_transaction,
		td.costumer_name,
		td.costumer_phone,
		td.seat
	FROM transaction_detail td
	JOIN transactions t ON td.id_transaction = t.id
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
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		details = append(details, detail)
	}

	return details, nil
}
