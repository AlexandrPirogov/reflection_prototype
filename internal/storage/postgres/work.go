package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/auth/user"
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/sheet"
)

func (pg *pgConnection) StartWork(u user.User, r sheet.SheetRow, p process.Process) error {
	query := `insert into work_sessions values(default, 
		(select id from users where email = $1),
		(select sc.id from sheets_content sc
			join sheets s on s.id = sc.sheets_id
			join processes p on p.id = s.proc_id and p.title = $2
			join users u on u.id = p.user_id and u.email = $1
			where sc.theme = $3), 
			now(), now())`

	_, err := pg.conn.Exec(context.Background(), query, u.Email, p.Title, r.Theme)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (pg *pgConnection) StopWork(u user.User, r sheet.SheetRow, p process.Process) error {
	query := `update work_sessions 
		set dt_end = now()
		where sheet_content_id =
		(select sc.id from sheets_content sc
			join sheets s on s.id = sc.sheets_id
			join processes p on p.id = s.proc_id and p.title = $2
			join users u on u.id = p.user_id and u.email = $1
			where sc.theme = $3)
		and dt_start = dt_end`

	_, err := pg.conn.Exec(context.Background(), query, u.Email, p.Title, r.Theme)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
