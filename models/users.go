package models

import (
	"context"
	"weeklytickits/utils"

	"github.com/jackc/pgx/v5"
)

type Users struct {
	Username     string `json:"username" form:"username"`
	Phone_number string `json:"phone" form:"phone"`
	Email        string `json:"email" form:"email"`
	Image        string `json:"image" form:"image"`
	Password     string `json:"password" form:"password"`
	Role         string `json:"role" form:"role"`
}

func FindAllUser() ([]Users, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return []Users{}, err
	}
	query := `SELECT * FROM users`
	data, err := conn.Query(context.Background(), query)
	result, err := pgx.CollectRows[Users](data, pgx.RowToStructByName)
	return result, nil
}
