package usecase

import (
	"context"
	"testing"
	"time"

	"github.com/mrizalr/go-tiket/domain"
	mocks "github.com/mrizalr/go-tiket/domain/mocks/user"
	"github.com/mrizalr/go-tiket/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecase_Create(t *testing.T) {
	userRepository := new(mocks.UserRepository)
	userUsecase := UserUsecase{userRepository}

	ctx := context.Background()
	req := model.UserRegisterRequest{
		Username: "test",
		Password: "test",
		Email:    "test@test.com",
		Gender:   1,
	}

	userRepository.Mock.On("Create", ctx, mock.AnythingOfType("model.UserRegisterRequest")).Return(domain.User{
		ID:        1,
		Username:  req.Username,
		Email:     req.Email,
		Gender:    1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil)

	user, err := userUsecase.Create(ctx, req)
	assert.NoError(t, err)

	userRepository.Mock.AssertExpectations(t)
	assert.NotZero(t, user.ID)
	assert.Equal(t, req.Username, user.Username)
	assert.Equal(t, req.Email, user.Email)
	assert.Equal(t, req.Gender, user.Gender)
	assert.NotEmpty(t, user.CreatedAt)
	assert.NotEmpty(t, user.UpdatedAt)
}
