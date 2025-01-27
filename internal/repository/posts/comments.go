package posts

import (
	"context"

	"github.com/Alam049/golang-campus/internal/model/posts"
)

func (r *respository) CreateComment(ctx context.Context, model posts.CommentModel) error {
	query := `INSERT INTO comments (user_id ,post_id, content, created_at, updated_at, created_by, updated_by) VALUES (? ,? ,? ,? ,? ,? ,?)`
	_, err := r.db.ExecContext(ctx, query, model.UserID, model.PostID, model.CommentContent, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}
