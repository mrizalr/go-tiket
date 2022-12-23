package mysql

import (
	"context"

	"github.com/mrizalr/go-tiket/domain"
	"github.com/mrizalr/go-tiket/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, userRegisterRequest model.UserRegisterRequest) (domain.User, error) {
	newUser := domain.User{
		Avatar:      userRegisterRequest.Avatar,
		Username:    userRegisterRequest.Username,
		Email:       userRegisterRequest.Email,
		PhoneNumber: userRegisterRequest.PhoneNumber,
		Gender:      userRegisterRequest.Gender,
		Birthdate:   userRegisterRequest.Birthdate,
	}

	result := r.db.Create(&newUser).Debug()
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	userCredential := domain.Credential{
		UserID:   newUser.ID,
		Password: userRegisterRequest.Password,
		User:     newUser,
	}

	result = r.db.Create(&userCredential).Debug()
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return newUser, nil
}

func (r *userRepository) GetCredential(ctx context.Context, id int) (domain.Credential, error) {
	credential := domain.Credential{}
	result := r.db.Where("id = ?", id).First(&credential)
	return credential, result.Error
}
