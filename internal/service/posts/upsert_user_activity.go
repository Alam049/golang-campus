package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/Alam049/golang-campus/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivity(ctx context.Context, userID, postID int64, req posts.UserActivitiesRequest) error {
	now := time.Now()
	model := posts.UserActivitiesModel{
		UserID:    userID,
		PostID:    postID,
		IsLiked:   req.IsLiked,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(userID, 10),
		UpdatedBy: strconv.FormatInt(userID, 10),
	}
	userActivity, err := s.postRepo.GetUserActivities(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("error get user activity from database")
		return err
	}
	if userActivity == nil {
		if !req.IsLiked {
			return errors.New("user activity not found")
		}
		err = s.postRepo.CreateUserActivities(ctx, model)
	} else {
		err = s.postRepo.UpdateUserActivities(ctx, model)
	}
	if err != nil {
		log.Error().Err(err).Msg("failed to upsert user activity to repository")
		return err
	}
	return nil
}
