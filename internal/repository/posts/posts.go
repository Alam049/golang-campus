package posts

import (
	"context"

	"github.com/Alam049/golang-campus/internal/model/posts"
)

func (r *respoitory) CreatePost(ctx context.Context, model posts.PostModel) error {
	query := `INSERT INTO posts (user_id ,post_title, post_content, hashtag, created_at, updated_at, created_by, updated_by) VALUES (? ,? ,? ,? ,? ,? ,? ,?)`
	_, err := r.db.ExecContext(ctx, query, model.UserID, model.PostTitle, model.PostContent, model.Hashtag, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}
