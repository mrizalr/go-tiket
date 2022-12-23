package domain

import (
	"context"
	"time"

	"github.com/mrizalr/go-tiket/model"
)

type User struct {
	ID          uint64    `json:"id" gorm:"primaryKey"`
	Avatar      string    `json:"avatar" gorm:"type:varchar(255)"`
	Username    string    `json:"username" gorm:"type:varchar(255);not null;unique"`
	Email       string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(255)"`
	Gender      uint8     `json:"gender" gorm:"not null;default:0"`
	Birthdate   time.Time `json:"birthdate"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type Credential struct {
	UserID   uint64 `json:"user_id" gorm:"primaryKey"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	User     User
}

type UserRepository interface {
	Create(context.Context, model.UserRegisterRequest) (User, error)
	GetCredential(context.Context, int) (Credential, error)
}

type UserUsecase interface {
	Create(context.Context, model.UserRegisterRequest) (User, error)
}
