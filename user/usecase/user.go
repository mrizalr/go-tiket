package usecase

import (
	"context"

	"github.com/mrizalr/go-tiket/domain"
	"github.com/mrizalr/go-tiket/model"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepository domain.UserRepository
}

func NewUserUsecase(userRepository domain.UserRepository) domain.UserUsecase {
	return &UserUsecase{
		userRepository,
	}
}

func (u *UserUsecase) Create(ctx context.Context, userRegisterRequest model.UserRegisterRequest) (domain.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(userRegisterRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, err
	}

	userRegisterRequest.Password = string(hashed)

	return u.userRepository.Create(ctx, userRegisterRequest)
}
