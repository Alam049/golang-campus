package memberships

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"

	"github.com/Alam049/golang-campus/internal/model/memberships"
	"github.com/Alam049/golang-campus/pkg/jwt"
)

func (s *service) Login(ctx context.Context, req memberships.LoginRequest) (string, error) {
	user, err := s.membershipRepo.GetUser(ctx, req.Email, "")
	if err != nil {
		log.Error().Err(err).Msg("Failed to get user")
		return "", err
	}
	if user == nil {
		return "", errors.New("email not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", errors.New("email or password is invalid ")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}
	return token, nil
}
