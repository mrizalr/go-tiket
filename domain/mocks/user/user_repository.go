package mocks

import (
	"context"

	"github.com/mrizalr/go-tiket/domain"
	"github.com/mrizalr/go-tiket/model"
	"github.com/stretchr/testify/mock"
)

type UserRepository struct {
	Mock mock.Mock
}

func (r *UserRepository) Create(ctx context.Context, req model.UserRegisterRequest) (domain.User, error) {
	args := r.Mock.Called(ctx, req)
	return args.Get(0).(domain.User), args.Error(1)
}

func (r *UserRepository) GetCredential(ctx context.Context, id int) (domain.Credential, error) {
	args := r.Mock.Called(ctx, id)
	return args.Get(0).(domain.Credential), args.Error(1)
}
