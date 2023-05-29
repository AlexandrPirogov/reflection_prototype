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

func (pg *pgConnection) ReadSheet(u user.User, p process.Process) (sheet.Sheet, error) {
	var res sheet.Sheet
	query := `select p.title, s.title from sheets s
				join processes p on s.proc_id = p.id and p.title = $1
				join users u on u.id = p.user_id and u.email = $2
				join sheets_content sc on s.id = sc.sheets_id`

	err := pg.conn.QueryRow(context.Background(), query, p.Title, u.Email).Scan(&res.Process, &res.Title)
	if err != nil {
		log.Println(err)
		return sheet.Sheet{}, err
	}

	contentQ := `select theme, date, done, sum(AGE(dt_end, dt_start))::varchar as spent from sheets_content sc
	join sheets s on s.id = sc.sheets_id
	join processes p on s.proc_id = p.id and p.title = $1
	join users u on u.id = p.user_id and u.email = $2
	join work_sessions ws on ws.sheet_content_id = sc.id
	group by theme, date, done`
	rows, err := pg.conn.Query(context.Background(), contentQ, p.Title, u.Email)

	if err != nil {
		log.Println(err)
		return sheet.Sheet{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var row sheet.SheetRow

		err = rows.Scan(&row.Theme, &row.Date, &row.Done, &row.Spent)

		if err != nil {
			continue
		}
		res = sheet.Add(row, res)
	}
	return res, nil
}

func (pg *pgConnection) AddRow(u user.User, r sheet.SheetRow, p process.Process) error {
	query := `insert into sheets_content values (default,
		(select id from sheets where proc_id = 
			(select p.id from processes p
				join users u on u.id = p.user_id and  p.title = $1 and u.email = $2)),
			$3, $4, $5)`

	_, err := pg.conn.Exec(context.Background(), query, p.Title, u.Email, r.Theme, r.Date, false)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (pg *pgConnection) MarkRow(u user.User, r sheet.SheetRow, p process.Process) error {
	query := `update sheets_content sc
	set done = true
	where sc.id = (select sc.id from sheets_content sc
		join sheets s on s.id = sc.sheets_id
		join processes p on p.id = s.proc_id and p.title = $1
		join users u on u.id = p.user_id and u.email = $2
		where sc.Theme = $3)
	`

	_, err := pg.conn.Exec(context.Background(), query, p.Title, u.Email, r.Theme)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
