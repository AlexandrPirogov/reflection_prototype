package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/config/env"

	"github.com/jackc/pgx/v5"
)

type pgConnection struct {
	conn *pgx.Conn
}

func New() *pgConnection {
	tmpUrl := env.ReadPgUrl()
	conn, err := pgx.Connect(context.Background(), tmpUrl)
	if err != nil {
		log.Fatal(err)
	}

	return &pgConnection{
		conn: conn,
	}
}
