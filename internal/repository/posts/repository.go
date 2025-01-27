package posts

import (
	"database/sql"
)

type respoitory struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *respoitory {
	return &respoitory{
		db: db,
	}
}
