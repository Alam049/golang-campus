package posts

import "time"

type (
	UserActivitiesRequest struct {
		IsLiked bool `json:"isLiked"`
	}
)

type (
	UserActivitiesModel struct {
		ID        int64     `db:"id"`
		UserID    int64     `db:"user_id"`
		PostID    int64     `db:"post_id"`
		IsLiked   bool      `db:"is_liked"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
		CreatedBy string    `db:"created_by"`
		UpdatedBy string    `db:"updated_by"`
	}
)
