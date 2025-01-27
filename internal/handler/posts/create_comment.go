package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/Alam049/golang-campus/internal/model/posts"
	"github.com/Alam049/golang-campus/internal/utils"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateComment(c *gin.Context) {
	ctx := c.Request.Context()

	var request posts.CreateCommentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		errorMsg := utils.ParseValidationError(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMsg,
		})
		return
	}

	userID := c.GetInt64("userID")
	postIDStr := c.Param("postID")
	postID, err := strconv.ParseInt(postIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errors.New("PostID not Valid").Error(),
		})
		return
	}

	err = h.postSvc.CreateComment(ctx, userID, postID, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}
