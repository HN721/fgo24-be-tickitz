package models

import (
	"context"
	"database/sql"
	"fmt"
	"weeklytickits/dto"
	"weeklytickits/utils"
)

func GetProfileByUserId(userId int) (dto.Profile, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return dto.Profile{}, err
	}
	defer conn.Conn().Close(context.Background())

	query := `SELECT id, fullname, phone_number, profile_image, id_user FROM profile WHERE id_user = $1`

	var profile dto.Profile
	var fullname, phone, image sql.NullString

	err = conn.QueryRow(context.Background(), query, userId).Scan(
		&profile.Id,
		&fullname,
		&phone,
		&image,
		&profile.IdUser,
	)
	if err != nil {
		return dto.Profile{}, err
	}

	profile.Fullname = nullStringToString(fullname)
	profile.Phone = nullStringToString(phone)
	profile.Image = nullStringToString(image)

	return profile, nil
}

func UpdateProfile(userId int, profile dto.Profile) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Conn().Close(context.Background())

	var fullname, phone, image sql.NullString

	queryGet := `SELECT fullname, phone_number, profile_image FROM profile WHERE id_user = $1`
	err = conn.QueryRow(context.Background(), queryGet, userId).Scan(
		&fullname,
		&phone,
		&image,
	)
	if err != nil {
		return fmt.Errorf("failed to get existing profile: %w", err)
	}

	if profile.Fullname == "" {
		profile.Fullname = nullStringToString(fullname)
	}
	if profile.Phone == "" {
		profile.Phone = nullStringToString(phone)
	}
	if profile.Image == "" {
		profile.Image = nullStringToString(image)
	}

	queryUpdate := `UPDATE profile SET fullname=$1, phone_number=$2, profile_image=$3 WHERE id_user=$4`
	_, err = conn.Exec(context.Background(), queryUpdate,
		profile.Fullname,
		profile.Phone,
		profile.Image,
		userId,
	)
	if err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}

	return nil
}

func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
