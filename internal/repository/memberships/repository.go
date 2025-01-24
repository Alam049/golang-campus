package memberships

import (
	"database/sql"
)

type respository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *respository {
	return &respository{
		db: db,
	}
}
