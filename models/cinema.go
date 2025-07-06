package models

import (
	"context"
	"fmt"
	"weeklytickits/utils"
)

type Cinema struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Logo string `json:"logo"`
}

func InsertCinema(cinema Cinema) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())
	query := `INSERT INTO cinema(name,logo)VALUES($1,$2)`
	_, err = conn.Exec(context.Background(), query,
		cinema.Name,
		cinema.Logo,
	)

	return err
}
func GetAllCinemas() ([]Cinema, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	rows, err := conn.Query(context.Background(), `SELECT id, name, logo FROM cinema`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cinemas []Cinema
	for rows.Next() {
		var c Cinema
		err := rows.Scan(&c.Id, &c.Name, &c.Logo)
		if err != nil {
			return nil, err
		}
		cinemas = append(cinemas, c)
	}

	return cinemas, nil
}

func GetCinemaByID(id int) (*Cinema, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Conn().Close(context.Background())

	var c Cinema
	err = conn.QueryRow(context.Background(),
		`SELECT id, name, logo FROM cinema WHERE id = $1`, id,
	).Scan(&c.Id, &c.Name, &c.Logo)

	if err != nil {
		return nil, fmt.Errorf("Cinema with ID %d not found: %v", id, err)
	}

	return &c, nil
}

func UpdateCinema(id int, cinema Cinema) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	var oldCinema Cinema
	queryGet := `SELECT name, logo FROM cinema WHERE id = $1`
	err = conn.QueryRow(context.Background(), queryGet, id).Scan(&oldCinema.Name, &oldCinema.Logo)
	if err != nil {
		return fmt.Errorf("failed to get existing cinema: %w", err)
	}

	if cinema.Name == "" {
		cinema.Name = oldCinema.Name
	}
	if cinema.Logo == "" {
		cinema.Logo = oldCinema.Logo
	}

	query := `UPDATE cinema SET name = $1, logo = $2 WHERE id = $3`
	_, err = conn.Exec(context.Background(), query, cinema.Name, cinema.Logo, id)
	if err != nil {
		return fmt.Errorf("failed to update cinema: %w", err)
	}

	return nil
}

func DeleteCinema(id int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	query := `DELETE FROM cinema WHERE id = $1`
	_, err = conn.Exec(context.Background(), query, id)
	return err
}
