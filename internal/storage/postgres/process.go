package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/process"
	"time"
)

// StoreProcess stores given process to db
//
// Pre-cond: given process to store. Process must be unique
//
// Post-cond: process was stored in db
func (pg *pgConnection) StoreProcess(p process.Process) error {
	_, err := pg.conn.Exec(context.Background(), "insert into processes values(default, 1, $1, $2)", process.Title(p), time.Now())
	if err != nil {
		log.Println(err)
		return err
	}

	if err = pg.contributeProcessCreation(); err != nil {
		return err
	}
	return nil
}

// ReadProcesses select processes with given pattern process
//
// Pre-cond: given pattern process
//
// Post-cond: returned  process that satisfied pattern
func (pg *pgConnection) ReadProcess(pattern process.Process) (process.Process, error) {
	var res process.Process
	query := "select title from processes where title = $1"
	err := pg.conn.QueryRow(context.Background(), query, pattern.Title).Scan(&res.Title)
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
func (pg *pgConnection) ListProcesses() ([]process.Process, error) {
	result := make([]process.Process, 0)
	rows, err := pg.conn.Query(context.Background(), "select title from processes")
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
