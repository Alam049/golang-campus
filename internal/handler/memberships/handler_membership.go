package memberships

import (
	"context"

	"github.com/Alam049/golang-campus/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface {
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
}

type Handler struct {
	*gin.Engine

	membershipSvc membershipService
}

func NewHandler(api *gin.Engine, membershipsSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSvc: membershipsSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("/memberships")
	route.GET("/ping", h.Ping)
	route.POST("/sign-up", h.SignUp)
}
