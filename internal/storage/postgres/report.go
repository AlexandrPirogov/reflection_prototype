package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/auth/user"
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/report"
	"reflection_prototype/internal/core/sheet"
)

const (
	QueryCreateReport = `insert into reports values(default,
		(select id from users u where u.email = $1),
		$2, now())`

	QueryFillReport = `insert into report_content values(
		default, 
		(select r.id from reports r 
		join users u on r.user_id = u.id and u.email = $1 and r.title = $2),
		(select sc.id from sheets_content sc
			join sheets s on s.id = sc.sheets_id
			join processes p on p.id = s.proc_id and p.title = $3
			join users u on u.id = p.user_id and u.email = $4
			where sc.Theme = $5)
		)`

	QueryReadReport = `select title from reports r
	join users u on u.id = r.user_id and u.email = $1 and r.title = $2`

	QueryReadReportsContent = `select sc.theme, sum(age(dt_end, dt_start))::varchar from report_content rc
	join reports r on r.id = rc.report_id and r.title = $1
	join users u on u.id = r.user_id and u.email = $2
	join sheets_content sc on rc.sc_id = sc.id
	join work_sessions ws on ws.sheet_content_id = sc.id
	group by sc.theme
	`
)

func (pg *pgConnection) CreateReport(u user.User, r report.Report) error {
	_, err := pg.conn.Exec(context.Background(), QueryCreateReport, u.Email, r.Title)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (pg *pgConnection) FillReport(u user.User, s sheet.SheetRow, p process.Process, r report.Report) error {
	_, err := pg.conn.Exec(context.Background(), QueryFillReport, u.Email, r.Title, p.Title, u.Email, s.Theme)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (pg *pgConnection) ReadReport(u user.User, r report.Report) (report.Report, error) {
	var res report.Report

	err := pg.conn.QueryRow(context.Background(), QueryReadReport, u.Email, r.Title).Scan(&res.Title)
	if err != nil {
		log.Println(err)
		return report.Report{}, err
	}

	rows, err := pg.conn.Query(context.Background(), QueryReadReportsContent, r.Title, u.Email)
	if err != nil {
		log.Println(err)
		return report.Report{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var sheetRow sheet.SheetRow

		err := rows.Scan(&sheetRow.Theme, &sheetRow.Spent)
		if err != nil {
			continue
		}

		res = report.Add(sheetRow, res)
	}

	return res, nil

}
