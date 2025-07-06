package models

import (
	"context"
	"fmt"
	"weeklytickits/dto"
	"weeklytickits/utils"

	"github.com/jackc/pgx/v5"
)

func GetProfileByUserId(userId int) (dto.Profile, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return dto.Profile{}, err
	}
	query := `SELECT id,fullname,phone_number,profile_image,id_user FROM profile WHERE id_user = $1`
	rows, err := conn.Query(context.Background(), query, userId)
	if err != nil {
		return dto.Profile{}, err
	}
	defer conn.Conn().Close(context.Background())

	data, err := pgx.CollectOneRow[dto.Profile](rows, pgx.RowToStructByName)
	return data, err

}
func UpdateProfile(userId int, profile dto.Profile) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	var oldProfile dto.Profile

	queryGet := `SELECT fullname, phone_number, profile_image FROM profile WHERE id_user = $1`
	err = conn.QueryRow(context.Background(), queryGet, userId).Scan(
		&oldProfile.Fullname,
		&oldProfile.Phone,
		&oldProfile.Image,
	)
	if err != nil {
		return fmt.Errorf("failed to get existing profile: %w", err)
	}

	if !profile.Fullname.Valid {
		profile.Fullname = oldProfile.Fullname
	}
	if !profile.Phone.Valid {
		profile.Phone = oldProfile.Phone
	}
	if !profile.Image.Valid {
		profile.Image = oldProfile.Image
	}

	queryUpdate := `UPDATE profile SET fullname=$1, phone_number=$2, profile_image=$3 WHERE id_user=$4`
	_, err = conn.Exec(context.Background(), queryUpdate, profile.Fullname, profile.Phone, profile.Image, userId)
	if err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}

	return nil
}
