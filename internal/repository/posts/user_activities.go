package posts

import (
	"context"
	"database/sql"

	"github.com/Alam049/golang-campus/internal/model/posts"
)

func (r *respoitory) GetUserActivities(ctx context.Context, model posts.UserActivitiesModel) (*posts.UserActivitiesModel, error) {
	query := `SELECT id, user_id, post_id, is_liked, created_at, updated_at, created_by, updated_by 
              FROM user_activities 
              WHERE user_id = ? AND post_id = ?`

	var response posts.UserActivitiesModel
	row := r.db.QueryRowContext(ctx, query, model.UserID, model.PostID)

	err := row.Scan(&response.ID, &response.UserID, &response.PostID, &response.IsLiked, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &response, nil
}

func (r *respoitory) CreateUserActivities(ctx context.Context, model posts.UserActivitiesModel) error {
	query := `INSERT INTO user_activities (user_id ,post_id, is_liked, created_at, updated_at, created_by, updated_by) VALUES (? ,? ,? ,? ,? ,? ,?)`
	_, err := r.db.ExecContext(ctx, query, model.UserID, model.PostID, model.IsLiked, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *respoitory) UpdateUserActivities(ctx context.Context, model posts.UserActivitiesModel) error {
	query := `UPDATE user_activities SET is_liked = ?, updated_at = ?, updated_by = ? WHERE user_id = ? AND post_id = ?`
	_, err := r.db.ExecContext(ctx, query, model.IsLiked, model.UpdatedAt, model.UpdatedBy, model.UserID, model.PostID)
	if err != nil {
		return err
	}
	return nil
}

func (r *respoitory) CountLikeByPostID(ctx context.Context, postID int64) (int, error) {
	query := `SELECT COUNT(id) FROM user_activities WHERE post_id = ? AND is_liked = true`

	var response int
	row := r.db.QueryRowContext(ctx, query, postID)

	err := row.Scan(&response)
	if err != nil {
		return response, err
	}
	return response, nil
}
