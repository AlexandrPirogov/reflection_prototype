package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/config/env"
	"reflection_prototype/internal/core/thread"
	"time"

	"github.com/jackc/pgx/v5"
)

// StoreThread stores given thread to db
//
// Pre-cond: given thread to store. Thread must be unique and process for thread must exists
//
// Post-cond: thread was stored in db
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

// ReadThreads select thread with given pattern thread
//
// Pre-cond: given pattern thread
//
// Post-cond: returned list of threads that satisfied pattern
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
