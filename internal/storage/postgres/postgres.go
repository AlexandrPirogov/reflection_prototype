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

	_, err = conn.Exec(context.Background(), "insert into processes values(default, 1, $1, $2)", process.Title(p), time.Now())
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
	rows, err := conn.Query(context.Background(), "select title from processes where title = $1", pattern.Title)
	if err != nil {
		log.Println(err)
		return nil, err
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

func (pg *PgConnection) StoreThread(t thread.Thread) error {
	tmpUrl := env.ReadPgUrl()
	conn, err := pgx.Connect(context.Background(), tmpUrl)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	query := `insert into threads values(default,
		(select id from processes where title = $1), $2, $3)`
	_, err = conn.Exec(context.Background(), query, t.Process, t.Title, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (pg *PgConnection) ReadThreads(t thread.Thread) ([]thread.Thread, error) {
	tmpUrl := env.ReadPgUrl()
	conn, err := pgx.Connect(context.Background(), tmpUrl)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	query := `select title, created_at from threads
	where title = $1 and proc_id = (select id from processes where title = $2)`

	rows, err := conn.Query(context.Background(), query, t.Title, t.Process)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := make([]thread.Thread, 0)

	defer rows.Close()
	for rows.Next() {
		var thread thread.Thread
		err := rows.Scan(&thread.Title, &thread.CreatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, thread)
	}
	return result, nil
}

func (pg *PgConnection) StoreQuant(q quant.Quant) error {
	tmpUrl := env.ReadPgUrl()
	conn, err := pgx.Connect(context.Background(), tmpUrl)
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())

	query := `insert into quants values(default,
		(select id from threads where title = $1
		and proc_id = (select id from processes where title = $2)), $3, $4, $5)`
	_, err = conn.Exec(context.Background(), query, q.Thread, q.Process, q.Title, q.Text, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (pg *PgConnection) ReadQuants(q quant.Quant) ([]quant.Quant, error) {
	tmpUrl := env.ReadPgUrl()
	conn, err := pgx.Connect(context.Background(), tmpUrl)
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())

	query := `select title, created_at from quants
	where title = $1 
	and thread_id = (select id from threads where title = $2
		and proc_id = (select id from processes where title = $3))`

	rows, err := conn.Query(context.Background(), query, q.Title, q.Thread, q.Process)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := make([]quant.Quant, 0)

	defer rows.Close()
	for rows.Next() {
		var q quant.Quant
		err := rows.Scan(&q.Title, &q.CreatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, q)
	}
	return result, nil
}
