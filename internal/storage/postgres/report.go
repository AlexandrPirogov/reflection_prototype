package storage

import (
	"context"
	"log"
	"reflection_prototype/internal/core/auth/user"
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/report"
	"reflection_prototype/internal/core/sheet"
)

/*
CREATE TABLE REPORTS(
    ID SERIAL PRIMARY KEY,
    USER_ID INT,
    TITLE VARCHAR(255),
    CREATED_AT TIMESTAMP,
    CONSTRAINT fk_reports_users FOREIGN KEY(USER_ID) REFERENCES USERS(ID)
);
CREATE UNIQUE INDEX reports_user_title ON REPORTS (USER_ID, TITLE);

CREATE TABLE REPORT_CONTENT(
    ID SERIAL PRIMARY KEY,
    REPORT_ID INT,
    SC_ID INT,
    CONSTRAINT fk_rc_report FOREIGN KEY(REPORT_ID) REFERENCES USERS(ID),
    CONSTRAINT fk_rc_sc FOREIGN KEY(SC_ID) REFERENCES SHEETS_CONTENT(ID)
);
*/

func (pg *pgConnection) CreateReport(u user.User, r report.Report) error {
	query := `insert into reports values(default,
		(select id from users u where u.email = $1),
		$2, now())`
	_, err := pg.conn.Exec(context.Background(), query, u.Email, r.Title)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (pg *pgConnection) FillReport(u user.User, s sheet.SheetRow, p process.Process, r report.Report) error {
	query := `insert into report_content values(
		default, 
		(select r.id from reports r 
		join users u on r.user_id = u.id and u.email = $1 and r.title = $2),
		(select sc.id from sheets_content sc
			join sheets s on s.id = sc.sheets_id
			join processes p on p.id = s.proc_id and p.title = $3
			join users u on u.id = p.user_id and u.email = $4
			where sc.Theme = $5)
		)`

	_, err := pg.conn.Exec(context.Background(), query, u.Email, r.Title, p.Title, u.Email, s.Theme)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (pg *pgConnection) ReadReport(u user.User, r report.Report) (report.Report, error) {
	var res report.Report

	query := `
	select title from reports r
	join users u on u.id = r.user_id and u.email = $1 and r.title = $2`

	err := pg.conn.QueryRow(context.Background(), query, u.Email, r.Title).Scan(&res.Title)
	if err != nil {
		log.Println(err)
		return report.Report{}, err
	}

	query = `
	select sc.theme, sum(age(dt_end, dt_start))::varchar from report_content rc
	join reports r on r.id = rc.report_id and r.title = $1
	join users u on u.id = r.user_id and u.email = $2
	join sheets_content sc on rc.sc_id = sc.id
	join work_sessions ws on ws.sheet_content_id = sc.id
	group by sc.theme
	`

	rows, err := pg.conn.Query(context.Background(), query, r.Title, u.Email)
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
