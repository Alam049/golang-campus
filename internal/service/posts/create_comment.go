package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/Alam049/golang-campus/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreateComment(ctx context.Context, userID, postID int64, req posts.CreateCommentRequest) error {
	now := time.Now()
	model := posts.CommentModel{
		UserID:         userID,
		PostID:         postID,
		CommentContent: req.CommentContent,
		CreatedAt:      now,
		UpdatedAt:      now,
		CreatedBy:      strconv.FormatInt(userID, 10),
		UpdatedBy:      strconv.FormatInt(userID, 10),
	}
	err := s.postRepo.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("failed to create comment to repository")
		return err
	}
	return nil
}
