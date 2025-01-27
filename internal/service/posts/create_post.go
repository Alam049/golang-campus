package posts

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/Alam049/golang-campus/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error {
	Hashtag := strings.Join(req.Hashtag, ",")

	now := time.Now()
	model := posts.PostModel{
		UserID:      userID,
		PostTitle:   req.PostTitle,
		PostContent: req.PostContent,
		Hashtag:     Hashtag,
		CreatedAt:   now,
		UpdatedAt:   now,
		CreatedBy:   strconv.FormatInt(userID, 10),
		UpdatedBy:   strconv.FormatInt(userID, 10),
	}
	err := s.postRepo.CreatePost(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("failed to create post to repository")
		return err
	}
	return nil
}
