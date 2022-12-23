package main

import (
	"context"
	"time"

	"github.com/mrizalr/go-tiket/db"
	"github.com/mrizalr/go-tiket/model"
	"github.com/mrizalr/go-tiket/user/repository/mysql"
	"github.com/mrizalr/go-tiket/user/usecase"
)

func main() {
	db := db.New()

	userRepository := mysql.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)

	userUsecase.Create(context.Background(), model.UserRegisterRequest{
		Username:  "test",
		Password:  "test",
		Email:     "test@gmail.com",
		Gender:    1,
		Birthdate: time.Date(1998, 5, 22, 0, 0, 0, 0, time.UTC),
	})
}
