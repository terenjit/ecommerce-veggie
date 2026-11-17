package service

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/terenjit/ecommerce-veggie/user-service/config"
	"github.com/terenjit/ecommerce-veggie/user-service/internal/adapter/repository"
	"github.com/terenjit/ecommerce-veggie/user-service/internal/core/domain/entity"
	"github.com/terenjit/ecommerce-veggie/user-service/utils/conv"
)

type UserServiceInterface interface {
	SignIn(ctx context.Context, req *entity.UserEntity) (*entity.UserEntity, string, error)
}

type UserService struct {
	userRepo   repository.UserRepositoryInterface
	cfg        *config.Config
	JwtService JwtServiceInterface
}

func NewUserService(userRepo repository.UserRepositoryInterface, cfg *config.Config, jwtService JwtServiceInterface) UserServiceInterface {
	return &UserService{userRepo: userRepo, cfg: cfg, JwtService: jwtService}
}

func (s *UserService) SignIn(ctx context.Context, req *entity.UserEntity) (*entity.UserEntity, string, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Error().Err(err).Msgf("error while fetching user by email %s", req.Email)
		return nil, "", err
	}

	if checkPass := conv.CheckPasswordHash(req.Password, user.Password); !checkPass {
		log.Error().Msgf("error while checking password for user %s", req.Email)
		err := errors.New("invalid password")
		return nil, "", err
	}

	token, err := s.JwtService.GenerateToken(user.ID)
	if err != nil {
		log.Error().Err(err).Msgf("error while generating token for user %s", user.ID)
		return nil, "", err
	}

	SessionData := map[string]interface{}{
		"user_id":    user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"logged_in":  true,
		"created_at": time.Now().String(),
		"token":      token,
	}

	redisConn := config.NewRedisClient()
	err = redisConn.HSet(ctx, token, SessionData).Err()
	if err != nil {
		log.Error().Err(err).Msgf("error while setting session data for user %s", user.ID)
		return nil, "", err
	}

	return user, token, nil
}
