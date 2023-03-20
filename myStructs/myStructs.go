package myStructs

import "time"

type User struct {
	UserId        string `json:"id" `
	Email         string `json:"email" binding:"required"`
	First_name    string `json:"first_name" binding:"required"`
	Middle_name   string `json:"middle_name" binding:"required"`
	Phone_number  string `json:"phone_number" binding:"required"`
	Firebase_id   string `json:"firebase_id,omitempty" `
	Password      string `json:"-" `
	Profile_photo string `json:"profile_photo,omitempty"`
}

type LoginData struct {
	Email       string `json:"email" binding:"required"`
	Password    string `json:"Password" binding:"required"`
	Firebase_id string `json:"firebase_id" binding:"required"`
}

type LocationUpdate struct {
	CurrentLatitude  string    `json:"current_latitude" binding:"required"`
	CurrentLongitude string    `json:"current_longitude" binding:"required"`
	UserId           string    `json:"user_id" binding:"required"`
	OriginLatitude   string    `json:"origin_latitude" binding:"required"`
	OriginLongitude  string    `json:"origin_longitude" binding:"required"`
	MaxDistance      string    `json:"max_distance" binding:"required"`
	LastUpdate       time.Time `json:"time" `
	Email            string    `json:"email" `
	FirstName        string    `json:"first_name" `
	MiddleName       string    `json:"middle_name" `
	PhoneNumber      string    `json:"phone_number"`
}
