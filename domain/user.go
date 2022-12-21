package domain

import "time"

type User struct {
	ID          uint64    `json:"id"`
	Avatar      string    `json:"avatar"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Gender      uint8     `json:"gender"`
	Birthdate   time.Time `json:"birthdate"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Credentials struct {
	UserID   uint64 `json:"user_id"`
	Password string `json:"password"`
	User     User
}

type UserRepository interface{}

type UserUsecase interface{}
