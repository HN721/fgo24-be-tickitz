package dto

type Profile struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname,omitempty"`
	Phone    string `json:"phoneNumber,omitempty" db:"phone_number"`
	Image    string `json:"image,omitempty" db:"profile_image"`
	IdUser   int    `json:"id_user" db:"id_user"`
}
type ProfileRequest struct {
	Fullname *string `json:"fullname"`
	Phone    *string `json:"phoneNumber"`
	Image    *string `json:"image"`
}
type ProfileResponse struct {
	Fullname *string `json:"fullname,omitempty" example:"Hosea"`
	Phone    *string `json:"phone,omitempty" example:"+628123456789"`
	Image    *string `json:"image,omitempty" example:"https://example.com/profile.jpg"`
}
