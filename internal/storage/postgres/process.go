package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/auth/user"
	"reflection_prototype/internal/core/process"
	"time"
)

const (
	QueryStoreProcess = `insert into processes values(default, 
		(select id from users where email = $1), $2, $3)`

	QueryReadProcess = `select title from processes p
	join users u on u.id = p.user_id and u.email = $1 and p.title = $2`

	QueryListProcesses = `select title from processes p
	join users u on u.id = p.user_id and u.email = $1`
)

// StoreProcess stores given process to db
//
// Pre-cond: given process to store. Process must be unique
//
// Post-cond: process was stored in db
func (pg *pgConnection) StoreProcess(u user.User, p process.Process) error {
	_, err := pg.conn.Exec(context.Background(), QueryStoreProcess, u.Email, process.Title(p), time.Now())
	if err != nil {
		log.Println(err)
		return err
	}

	if err = pg.contributeProcessCreation(u); err != nil {
		return err
	}
	return nil
}

// ReadProcesses select processes with given pattern process
//
// Pre-cond: given pattern process
//
// Post-cond: returned  process that satisfied pattern
func (pg *pgConnection) ReadProcess(u user.User, pattern process.Process) (process.Process, error) {
	var res process.Process
	err := pg.conn.QueryRow(context.Background(), QueryReadProcess, u.Email, pattern.Title).Scan(&res.Title)
	if err != nil {
		log.Println(err)
		return process.Process{}, err
	}

	return res, nil
}

// ListProcesses select processes
//
// Pre-cond:
//
// Post-cond: returned  list of processes
func (pg *pgConnection) ListProcesses(u user.User) ([]process.Process, error) {
	result := make([]process.Process, 0)
	rows, err := pg.conn.Query(context.Background(), QueryListProcesses, u.Email)
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
