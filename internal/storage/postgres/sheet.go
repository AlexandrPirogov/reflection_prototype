package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/sheet"
)

func (pg *pgConnection) StoreSheet(s sheet.Sheet, p process.Process) error {
	query := `insert into sheets values(default, $1, 
		(select id from processes where title = $2))`

	_, err := pg.conn.Exec(context.Background(), query, s.Title, p.Title)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// TODO add content scan
func (pg *pgConnection) ReadSheet(p process.Process) (sheet.Sheet, error) {
	var res sheet.Sheet
	query := `select p.title, s.title from sheets s
				join processes p on s.proc_id = p.id and p.title = $1`

	err := pg.conn.QueryRow(context.Background(), query, p.Title).Scan(&res.Process, &res.Title)
	if err != nil {
		log.Println(err)
		return sheet.Sheet{}, err
	}

	return res, nil
}

func (pg *pgConnection) AddRow(r sheet.SheetRow, s sheet.Sheet) error {
	query := `insert into sheets_content values (default,
		(select distinct id from sheets where proc_id = 
			(select id from processes where title = $1)),
			$2, $3, $4, null)`

	_, err := pg.conn.Exec(context.Background(), query, s.Process, r.Theme, r.Date, false)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
