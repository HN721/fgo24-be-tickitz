package models

import (
	"context"
	"errors"
	"fmt"
	"weeklytickits/services"
	"weeklytickits/utils"

	"github.com/jackc/pgx/v5"
)

type Users struct {
	UserID       int    `json:"userId" db:"user_id"`
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
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	query := `SELECT * FROM users`
	data, err := conn.Query(context.Background(), query)
	result, err := pgx.CollectRows[Users](data, pgx.RowToStructByName)
	return result, nil
}
func Register(user Users) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	query := `
	INSERT INTO users (username, email, password)
	VALUES ($1, $2, $3)
`

	_, err = conn.Exec(context.Background(), query,
		user.Username,
		user.Email,
		user.Password,
	)
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	return err
}
func Login(user Users) (Users, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return user, err
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	var dbUser Users
	query := `SELECT user_id, username, email, password FROM users WHERE email = $1`
	err = conn.QueryRow(context.Background(), query, user.Email).
		Scan(&dbUser.UserID, &dbUser.Username, &dbUser.Email, &dbUser.Password)

	if err != nil {
		return Users{}, fmt.Errorf("email tidak ditemukan")
	}

	if user.Password != dbUser.Password {
		return Users{}, fmt.Errorf("password salah")
	}
	return dbUser, nil
}

func ChangePassword(userId int, newPassword string, oldPassword string) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	if newPassword == oldPassword {
		return err
	}
	query := `UPDATE users SET password = $1 WHERE user_id = $2`
	result, err := conn.Exec(context.Background(), query, newPassword, userId)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("user tidak ditemukan")
	}

	return nil
}
func ForgetPassword(email string) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	var foundEmail string
	query := `SELECT email FROM users WHERE email = $1`
	err = conn.QueryRow(context.Background(), query, email).Scan(&foundEmail)
	if err != nil {
		return fmt.Errorf("email not found: %w", err)
	}
	otp := services.GenerateOTP()

	subject := "Yout OTP Reset PASSWORD"
	body := fmt.Sprintf("<p>Your OTP is: <b>%s</b> </p>", otp)
	err = services.SendEmail(email, subject, body)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
func SaveOTP(otp string) {
	// save ke redis
}
