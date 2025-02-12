package posts

import "time"

type (
	CreateCommentRequest struct {
		CommentContent string `json:"commentContent" binding:"required"`
	}
)

type (
	CommentModel struct {
		ID             int64     `db:"id"`
		UserID         int64     `db:"user_id"`
		PostID         int64     `db:"post_id"`
		CommentContent string    `db:"content"`
		CreatedAt      time.Time `db:"created_at"`
		UpdatedAt      time.Time `db:"updated_at"`
		CreatedBy      string    `db:"created_by"`
		UpdatedBy      string    `db:"updated_by"`
	}
)
