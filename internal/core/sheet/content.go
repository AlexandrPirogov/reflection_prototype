package sheet

type sheetContent struct {
	Content []SheetRow `json:"Content"`
}

// newContent creates new Instance of sheetContent
func newContent() sheetContent {
	return sheetContent{
		make([]SheetRow, 0),
	}
}

// addRow adding row to sheetContent
//
// Pre-cond: given row to add and sheetContent to which add row
//
// Post-cond: row was added to sheetContent
func addRow(row SheetRow, sheetContent sheetContent) sheetContent {
	sheetContent.Content = append(sheetContent.Content, row)
	return sheetContent
}
