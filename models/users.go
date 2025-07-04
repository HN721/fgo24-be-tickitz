package models

import (
	"context"
	"fmt"
	"strconv"
	"time"
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
	defer func() {
		conn.Conn().Close(context.Background())
	}()

	tx, err := conn.Begin(context.Background())
	if err != nil {
		return err
	}
	defer tx.Rollback(context.Background())

	hash, err := services.HashPassword(user.Password)
	if err != nil {
		return err
	}

	var userId int
	err = tx.QueryRow(context.Background(), `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`, user.Username, user.Email, hash).Scan(&userId)
	if err != nil {
		return err
	}

	_, err = tx.Exec(context.Background(), `
		INSERT INTO profile (id_user)
		VALUES ($1)
	`, userId)
	if err != nil {
		return err
	}

	err = tx.Commit(context.Background())
	return err
}

func Login(user Users) (Users, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return Users{}, err
	}
	defer conn.Conn().Close(context.Background())

	var dbUser Users
	query := `SELECT username, email, password, role FROM users WHERE email = $1`
	err = conn.QueryRow(context.Background(), query, user.Email).
		Scan(&dbUser.Username, &dbUser.Email, &dbUser.Password, &dbUser.Role)

	if err != nil {
		if err.Error() == "no rows in result set" {
			return Users{}, fmt.Errorf("email tidak ditemukan")
		}
		return Users{}, err
	}

	if err := services.ComparePassword(dbUser.Password, user.Password); err != nil {
		return Users{}, fmt.Errorf("password salah")
	}

	return dbUser, nil
}

func ChangePassword(email string, OTP int, newPassword string, oldPassword string) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer func() {
		conn.Conn().Close(context.Background())
	}()
	var hashedPassword string
	query := `SELECT password FROM users WHERE email = $1`
	err = conn.QueryRow(context.Background(), query, email).Scan(&hashedPassword)
	if err != nil {
		return fmt.Errorf("user tidak ditemukan")
	}
	client := utils.RedisClient
	key := fmt.Sprintf("otp:%s", email)
	storedOTP, err := client.Get(context.Background(), key).Result()
	oldOtp, _ := strconv.Atoi(storedOTP)
	if oldOtp != OTP {
		return fmt.Errorf("Invalid OTP")
	}
	err = services.ComparePassword(hashedPassword, oldPassword)
	if err != nil {
		return fmt.Errorf("password lama salah")
	}

	if oldPassword == newPassword {
		return fmt.Errorf("password baru tidak boleh sama dengan password lama")
	}

	newHashed, err := services.HashPassword(newPassword)
	if err != nil {
		return fmt.Errorf("gagal hash password baru: %v", err)
	}

	updateQuery := `UPDATE users SET password = $1 WHERE email = $2`
	result, err := conn.Exec(context.Background(), updateQuery, newHashed, email)
	if err != nil {
		return fmt.Errorf("gagal update password: %v", err)
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("user tidak ditemukan")
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
	err = SaveOTP(email, otp)
	if err != nil {
		return err
	}
	subject := "Yout OTP Reset PASSWORD"
	body := fmt.Sprintf("<p>Your OTP is: <b>%d</b> </p>", otp)
	err = services.SendEmail(email, subject, body)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
func SaveOTP(email string, otp int) error {
	client := utils.RedisClient
	key := fmt.Sprintf("otp:%s", email)
	err := client.Set(context.Background(), key, otp, 5*time.Minute).Err()
	if err != nil {
		return fmt.Errorf("failed to save OTP to Redis: %w", err)
	}
	return nil
}
