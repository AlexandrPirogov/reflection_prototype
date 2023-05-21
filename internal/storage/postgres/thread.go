package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/thread"
	"time"
)

// StoreThread stores given thread to db
//
// Pre-cond: given thread to store. Thread must be unique and process for thread must exists
//
// Post-cond: thread was stored in db
func (pg *pgConnection) StoreThread(t thread.Thread) error {
	query := `insert into threads values(default,
		(select id from processes where title = $1), $2, $3)`
	_, err := pg.conn.Exec(context.Background(), query, t.Process, t.Title, time.Now())
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
func (pg *pgConnection) ReadThreads(t thread.Thread) ([]thread.Thread, error) {
	query := `select title, created_at from threads
	where title = $1 and proc_id = (select id from processes where title = $2)`

	rows, err := pg.conn.Query(context.Background(), query, t.Title, t.Process)
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
