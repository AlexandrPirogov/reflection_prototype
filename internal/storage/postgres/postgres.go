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

func (pg *PgConnection) ReadProcesses(pattern process.Process) ([]process.Process, error) {
	tmpUrl := env.ReadPgUrl()
	conn, err := pgx.Connect(context.Background(), tmpUrl)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	result := make([]process.Process, 0)
	rows, err := conn.Query(context.Background(), "select title from processes where title = $1", process.Title(pattern))
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var process process.Process
		err := rows.Scan(&process.Title)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, process)
	}
	return result, nil
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
