package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/auth/user"
	"reflection_prototype/internal/core/quant"
	"time"
)

// StoreQuant stores given quant to db
//
// Pre-cond: given quant to store. Quant must be unique and process and thread for quant must exist
//
// Post-cond: quant was stored in db
func (pg *pgConnection) StoreQuant(u user.User, q quant.Quant) error {
	query := `insert into quants values(default,
		(select threads.id from threads where title = $1
		and proc_id = (select p.id from processes p 
			join users u on u.id = p.user_id where title = $2 and u.email = $3)), $4, $5, $6)`
	_, err := pg.conn.Exec(context.Background(), query, q.Thread, q.Process, u.Email, q.Title, q.Text, time.Now())
	if err != nil {
		log.Println(err)
		return err
	}

	if err = pg.contributeQantCreation(); err != nil {
		return err
	}

	return nil
}

// ListQuants returns list of quants stored in db
//
// Pre-cond:
//
// Post-cond: returned list of quants
func (pg *pgConnection) ListQuants(u user.User) ([]quant.Quant, error) {
	query := `select p.title, t.title, q.title, q.text, q.created_at from quants q
			join threads t on q.thread_id = t.id
			join processes p on t.proc_id = p.id
			join users u on u.id = p.user_id and u.email = $1`

	rows, err := pg.conn.Query(context.Background(), query, u.Email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result := make([]quant.Quant, 0)

	defer rows.Close()
	for rows.Next() {
		var q quant.Quant
		err := rows.Scan(&q.Process, &q.Thread, &q.Title, &q.Text, &q.CreatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		result = append(result, q)
	}
	return result, nil
}

// ReadQuant return quants that stored in db and satisfied the given pattern
//
// Pre-cond: given pattern q Quant
//
// Post-cond: returned quant that satisfied given pattern
func (pg *pgConnection) ReadQuant(u user.User, q quant.Quant) (quant.Quant, error) {
	var res quant.Quant
	query := `select title, created_at from quants
	where title = $1 
	and thread_id = (select id from threads where title = $2
		and proc_id = (select id from processes p where title = $3
			join users u on u.id = p.user_id and u.email = $4))`

	err := pg.conn.QueryRow(context.Background(), query, q.Title, q.Thread, q.Process, u.Email).Scan(&res.Title, &res.CreatedAt)
	if err != nil {
		log.Println(err)
		return quant.Quant{}, err
	}

	return res, nil
}
