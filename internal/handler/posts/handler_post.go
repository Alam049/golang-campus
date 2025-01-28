package posts

import (
	"context"

	"github.com/Alam049/golang-campus/internal/middleware"
	"github.com/Alam049/golang-campus/internal/model/posts"
	"github.com/gin-gonic/gin"
)

type postService interface {
	CreatePost(ctx context.Context, userID int64, req posts.CreatePostRequest) error
	CreateComment(ctx context.Context, userID, postID int64, req posts.CreateCommentRequest) error
	UpsertUserActivity(ctx context.Context, userID, postID int64, req posts.UserActivitiesRequest) error
	GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error)
}

type Handler struct {
	*gin.Engine

	postSvc postService
}

func NewHandler(api *gin.Engine, postsSvc postService) *Handler {
	return &Handler{
		Engine:  api,
		postSvc: postsSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create", h.CreatePost)
	route.POST("/comment/:postID", h.CreateComment)
	route.PUT("/user_activity/:postID", h.UpsertUserActivity)
	route.GET("/list", h.GetAllPost)
}
