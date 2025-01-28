package posts

import (
	"context"

	"github.com/Alam049/golang-campus/internal/model/posts"
)

func (r *respoitory) CreateComment(ctx context.Context, model posts.CommentModel) error {
	query := `INSERT INTO comments (user_id ,post_id, content, created_at, updated_at, created_by, updated_by) VALUES (? ,? ,? ,? ,? ,? ,?)`
	_, err := r.db.ExecContext(ctx, query, model.UserID, model.PostID, model.CommentContent, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *respoitory) GetCommentsByPostID(ctx context.Context, postID int64) ([]posts.Comment, error) {
	query := `SELECT c.id, c.user_id, u.username, c.content FROM comments c JOIN users u ON c.user_id = u.id WHERE c.post_id = ? ORDER BY c.updated_at DESC`

	rows, err := r.db.QueryContext(ctx, query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := make([]posts.Comment, 0)
	for rows.Next() {
		var (
			comment  posts.Comment
			username string
		)
		err = rows.Scan(&comment.ID, &comment.UserID, &username, &comment.CommentContent)
		if err != nil {
			return response, err
		}
		response = append(response, posts.Comment{
			ID:             comment.ID,
			UserID:         comment.UserID,
			Username:       username,
			CommentContent: comment.CommentContent,
		})
	}
	return response, nil
}
