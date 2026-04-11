package main

import (
	"log"
	"net/http"

	"github.com/prantoran/goreddit/postgres"
	"github.com/prantoran/goreddit/web"
)

func main() {
	dsn := "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable"
	store, err := postgres.NewStore(dsn)
	if err != nil {
		log.Fatal(err)
	}

	sessions, err := web.NewSessionHandler(dsn)
	if err != nil {
		log.Fatal(err)
	}

	csrfKey := []byte("01234567890123456789012345678901")
	h := web.NewHandler(store, sessions, csrfKey)
	http.ListenAndServe(":3000", h)
}
