package memberships

import (
	"context"
	"errors"
	"time"

	"github.com/Alam049/golang-campus/internal/model/memberships"
	"github.com/Alam049/golang-campus/pkg/jwt"
	"github.com/rs/zerolog/log"
)

func (s *service) ValidateRefreshToken(ctx context.Context, userID int64, req memberships.RefreshTokenRequest) (string, error) {
	existingRefreshToken, err := s.membershipRepo.GetRefreshToken(ctx, userID, time.Now())

	if err != nil {
		log.Error().Err(err).Msg("failed to get refresh token")
		return "", err
	}

	if existingRefreshToken == nil {
		return "", errors.New("refresh token has been expired")
	}
	if existingRefreshToken.RefreshToken != req.Token {
		return "", errors.New("invalid refresh token")
	}

	user, err := s.membershipRepo.GetUser(ctx, "", "", userID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to ge  ")
		return "", err
	}
	if user == nil {
		return "", errors.New("user not exist")
	}

	token, err := jwt.CreateToken(user.ID, user.Username, s.cfg.Service.SecretJWT)
	if err != nil {
		return "", err
	}
	return token, nil
}
