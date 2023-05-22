package sheet

import "time"

type SheetRow struct {
	Theme string
	Date  time.Time
	Done  bool
}

// NewSheetRow creates instance of SheetRow
//
// Pre-cond: given Theme, Date of creation
//
// Post-cond: SheetRow instance was created
func NewSheetRow(Theme string, Date time.Time) SheetRow {
	return SheetRow{
		Theme: Theme,
		Date:  Date,
		Done:  false,
	}
}
