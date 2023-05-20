package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/config/env"
	"reflection_prototype/internal/core/process"
	"time"

	"github.com/jackc/pgx/v5"
)

// StoreProcess stores given process to db
//
// Pre-cond: given process to store. Process must be unique
//
// Post-cond: process was stored in db
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

// ReadProcesses select processes with given pattern process
//
// Pre-cond: given pattern process
//
// Post-cond: returned list of processes that satisfied pattern
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
