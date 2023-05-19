package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/config/env"
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/quant"
	"reflection_prototype/internal/core/thread"
	"time"

	"github.com/jackc/pgx/v5"
)

type PgConnection struct {
}

func (pg *PgConnection) StoreProcess(p process.Process) error {
	tmpUrl := env.ReadPgUrl()
	conn, err := pgx.Connect(context.Background(), tmpUrl)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	_, err = conn.Query(context.Background(), "insert into processes values(default, 1, $1, $2)", process.Title(p), time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (pg *PgConnection) SelectProcess(p process.Process) (process.Process, error) {
	return p, nil
}

func (pg *PgConnection) StoreThread(t thread.Thread, p process.Process) error {
	return nil
}
func (pg *PgConnection) SelectThread(t thread.Thread, p process.Process) (thread.Thread, error) {
	return t, nil
}

func (pg *PgConnection) StoreQuant(q quant.Quant, t thread.Thread, p process.Process) error {
	return nil
}
