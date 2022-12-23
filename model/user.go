package model

import "time"

type UserRegisterRequest struct {
	Avatar      string    `json:"avatar"`
	Username    string    `json:"username,omitempty"`
	Password    string    `json:"password,omitempty"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Gender      uint8     `json:"gender,omitempty"`
	Birthdate   time.Time `json:"birthdate"`
}
