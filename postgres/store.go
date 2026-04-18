package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // import postgres driver as side effect
	"github.com/prantoran/goreddit"
)

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Store{
		ThreadStore:  &ThreadStore{DB: db},
		PostStore:    &PostStore{DB: db},
		CommentStore: &CommentStore{DB: db},
		UserStore:    &UserStore{DB: db},
	}, nil
}

type Store struct {
	goreddit.ThreadStore
	goreddit.PostStore
	goreddit.CommentStore
	goreddit.UserStore
}
