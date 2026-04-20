package web

import (
	"context"
	"database/sql"
	"encoding/gob"

	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

func init() {
	gob.Register(uuid.UUID{})
}

func NewSessionHandler(dataSourceName string) (*scs.SessionManager, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	sessions := scs.New()
	sessions.Store = postgresstore.New(db)

	return sessions, nil
}

type SessionData struct {
	FlashMessage string
	Form         interface{}
	// UserID       uuid.UUID
}

func GetSessionData(session *scs.SessionManager, ctx context.Context) SessionData {
	var data SessionData
	data.FlashMessage = session.PopString(ctx, "flash")
	// data.UserID = session.Get(ctx, "user_id").(uuid.UUID)

	data.Form = session.Pop(ctx, "form")
	if data.Form == nil {
		data.Form = map[string]string{}
	}

	return data
}
