package posts

import (
	"context"

	"github.com/Alam049/golang-campus/internal/configs"
	"github.com/Alam049/golang-campus/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error
	GetUserActivities(ctx context.Context, model posts.UserActivitiesModel) (*posts.UserActivitiesModel, error)
	CreateUserActivities(ctx context.Context, model posts.UserActivitiesModel) error
	UpdateUserActivities(ctx context.Context, model posts.UserActivitiesModel) error
}

type service struct {
	cfg      *configs.Config
	postRepo postRepository
}

func NewService(cfg *configs.Config, postRepo postRepository) *service {
	return &service{
		cfg:      cfg,
		postRepo: postRepo,
	}
}
