package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/quant"
	"time"
)

// StoreQuant stores given quant to db
//
// Pre-cond: given quant to store. Quant must be unique and process and thread for quant must exist
//
// Post-cond: quant was stored in db
func (pg *pgConnection) StoreQuant(q quant.Quant) error {
	query := `insert into quants values(default,
		(select id from threads where title = $1
		and proc_id = (select id from processes where title = $2)), $3, $4, $5)`
	_, err := pg.conn.Exec(context.Background(), query, q.Thread, q.Process, q.Title, q.Text, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}

	if err = pg.contributeQantCreation(); err != nil {
		return err
	}

	return nil
}

// ReadQuants select quant with given pattern quant
//
// Pre-cond: given pattern quant
//
// Post-cond: returned list of quants that satisfied pattern
func (pg *pgConnection) ReadQuant(q quant.Quant) ([]quant.Quant, error) {
	query := `select title, created_at from quants
	where title = $1 
	and thread_id = (select id from threads where title = $2
		and proc_id = (select id from processes where title = $3))`

	rows, err := pg.conn.Query(context.Background(), query, q.Title, q.Thread, q.Process)
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