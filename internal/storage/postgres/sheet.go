package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/auth/user"
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/sheet"
)

func (pg *pgConnection) StoreSheet(u user.User, s sheet.Sheet, p process.Process) error {
	query := `insert into sheets values(default, $1, 
		(select p.id from processes p
			join users u on u.id = p.user_id where title = $2 and u.email = $3))`

	_, err := pg.conn.Exec(context.Background(), query, s.Title, p.Title, u.Email)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// TODO add content scan
func (pg *pgConnection) ReadSheet(u user.User, p process.Process) (sheet.Sheet, error) {
	var res sheet.Sheet
	query := `select p.title, s.title from sheets s
				join processes p on s.proc_id = p.id and p.title = $1
				join users u on u.id = p.user_id and u.email = $2`

	err := pg.conn.QueryRow(context.Background(), query, p.Title, u.Email).Scan(&res.Process, &res.Title)
	if err != nil {
		log.Println(err)
		return sheet.Sheet{}, err
	}

	return res, nil
}

func (pg *pgConnection) AddRow(u user.User, r sheet.SheetRow, s sheet.Sheet) error {
	query := `insert into sheets_content values (default,
		(select distinct id from sheets where proc_id = 
			(select id from processes p
				join user u on u.id = p.user_id and  p.title = $1 and u.email = $2)),
			$3, $4, $5, null)`

	_, err := pg.conn.Exec(context.Background(), query, s.Process, u.Email, r.Theme, r.Date, false)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
