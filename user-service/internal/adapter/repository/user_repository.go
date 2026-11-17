package repository

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"github.com/terenjit/ecommerce-veggie/user-service/internal/core/domain/entity"
	"github.com/terenjit/ecommerce-veggie/user-service/internal/core/domain/model"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
}
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	var user model.User
	if err := r.db.WithContext(ctx).Preload("Role").Where("email = ? AND is_verified = ?", email, true).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("404")
			log.Error().Err(err).Msgf("GetUserByEmail: user not found for email: %v", email)
			return nil, err
		}
		log.Error().Err(err).Msgf("GetUserByEmail: Failed to get user : %v", err)
		return nil, err
	}

	userEntity := &entity.UserEntity{
		ID:         user.ID,
		Email:      user.Email,
		Name:       user.Name,
		Password:   user.Password,
		RoleName:   user.Role.Name,
		IsVerified: user.IsVerified,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		Address:    user.Address,
		Phone:      user.Phone,
		Photo:      user.Photo,
		Lat:        user.Lat,
		Lng:        user.Lng,
	}

	return userEntity, nil
}
