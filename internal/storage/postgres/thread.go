package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/auth/user"
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/thread"
	"time"
)

// StoreThread stores given thread to db
//
// Pre-cond: given thread to store. Thread must be unique and process for thread must exists
//
// Post-cond: thread was stored in db
func (pg *pgConnection) StoreThread(u user.User, t thread.Thread) error {
	query := `insert into threads values(default,
		(select p.id from processes p 
			join users u on u.id = p.user_id and title = $1 and u.email = $2), $3, $4)`
	_, err := pg.conn.Exec(context.Background(), query, t.Process, u.Email, t.Title, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}

	if err = pg.contributeThreadCreation(); err != nil {
		return err
	}
	return nil
}

// ListThreads returns threads that stored in storage
//
// Pre-cond:
//
// Post-cond: returned list of threads that stored in storage
func (pg *pgConnection) ListThreads(u user.User) ([]thread.Thread, error) {
	query := `select p.title, t.title, t.created_at from threads t
				join processes p on p.id = t.proc_id
				join users u on u.id = p.user_id and u.email = $1`

	rows, err := pg.conn.Query(context.Background(), query, u.Email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := make([]thread.Thread, 0)

	defer rows.Close()
	for rows.Next() {
		var thread thread.Thread
		err := rows.Scan(&thread.Process, &thread.Title, &thread.CreatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, thread)
	}
	return result, nil
}

// ReadThreads select thread with given pattern thread
//
// Pre-cond: given pattern thread
//
// Post-cond: returned list of threads that satisfied pattern
func (pg *pgConnection) ReadThread(u user.User, t thread.Thread) (thread.Thread, error) {

	var result thread.Thread
	query := `select title, created_at from threads
	where title = $1 and proc_id = 
	(select id from processes p 
		join users u on u.id = p.user_id and title = $2 and u.email = $3)`

	err := pg.conn.QueryRow(context.Background(), query, t.Title, t.Process, u.Email).Scan(&result.Title, &result.CreatedAt)
	if err != nil {
		log.Println(err)
		return thread.Thread{}, err
	}

	return result, nil
}

// ListProcessesThreads select threads that belong to given process
//
// Pre-cond: given pattern process to which threads are belong
//
// Post-cond: returned list of threads that belong to given process
func (pg *pgConnection) ListProcessesThreads(u user.User, p process.Process) ([]thread.Thread, error) {
	query := `select p.title, t.title, t.created_at from threads t
	join processes p on p.id = t.proc_id and p.title = $1
	join users u on u.id = p.user_id and u.email = $2`

	rows, err := pg.conn.Query(context.Background(), query, p.Title, u.Email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := make([]thread.Thread, 0)

	defer rows.Close()
	for rows.Next() {
		var thread thread.Thread
		err := rows.Scan(&thread.Process, &thread.Title, &thread.CreatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, thread)
	}
	return result, nil
}
