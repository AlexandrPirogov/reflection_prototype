package storage

import (
	"log"
	"reflection_prototype/internal/core"
	"reflection_prototype/internal/core/sheet"

	"github.com/jackc/pgx/v5"
)

var rowsErrors = map[bool]error{
	false: pgx.ErrNoRows,
	true:  nil,
}

var pgSheetFactory = map[bool]func(process, title string) core.Sheeter{
	true: func(process, title string) core.Sheeter {
		return sheet.EmptySheet{
			Process: process,
			Title:   title,
		}
	},
	false: func(process, title string) core.Sheeter {
		return sheet.Sheet{
			Process: process,
			Title:   title,
		}
	},
}

var pgSheetErrors = map[error]map[core.CoreType]func(rows pgx.Rows) (core.Sheeter, error){
	pgx.ErrNoRows: {
		core.SHEET: func(rows pgx.Rows) (core.Sheeter, error) {
			log.Println("empty sheet")
			return sheet.EmptySheet{}, pgx.ErrNoRows
		},
	},
	nil: {
		core.SHEET: func(rows pgx.Rows) (core.Sheeter, error) {
			log.Println("non empty sheet")
			res := sheet.Sheet{}

			for rows.Next() {
				var row sheet.SheetRow
				err := rows.Scan(&row.Theme, &row.Date, &row.Done, &row.Spent)
				if err != nil {
					continue
				}
				res = sheet.Add(row, res)
			}
			rows.Close()
			return res, nil
		},
	},
}
