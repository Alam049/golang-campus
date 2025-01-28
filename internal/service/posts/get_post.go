package posts

import (
	"context"

	"github.com/Alam049/golang-campus/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context, postID int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostByID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get post by id")
		return nil, err
	}
	likeCount, err := s.postRepo.CountLikeByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to COUNT LIKE by id")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentsByPostID(ctx, postID)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get comments by post id")
		return nil, err
	}
	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID:          postDetail.ID,
			UserID:      postDetail.UserID,
			Username:    postDetail.Username,
			PostTitle:   postDetail.PostTitle,
			PostContent: postDetail.PostContent,
			Hashtag:     postDetail.Hashtag,
			IsLiked:     postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comments:  comments,
	}, nil
}
