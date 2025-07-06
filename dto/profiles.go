package dto

import "database/sql"

type Profile struct {
	Id       int            `json:"id"`
	Fullname sql.NullString `json:"fullname,omitempty"`
	Phone    sql.NullString `json:"phoneNumber,omitempty" db:"phone_number"`
	Image    sql.NullString `json:"image,omitempty" db:"profile_image"`
	IdUser   int            `json:"id_user" db:"id_user"`
}
type ProfileRequest struct {
	Fullname *string `json:"fullname"`
	Phone    *string `json:"phoneNumber"`
	Image    *string `json:"image"`
}
