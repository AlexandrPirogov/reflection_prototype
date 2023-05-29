/*
package sheet responsible for creating plans for processes like google sheets

relation between sheets and processes is 1-to-1
*/
package sheet

import "reflection_prototype/internal/core"

type Sheet struct {
	Process string
	Title   string
	Content sheetContent
}

func (s Sheet) IsEmpty() bool {
	return s.Content.IsEmpty()
}

func (e Sheet) SetTitle(Title string) core.Sheeter {
	e.Title = Title
	return e
}

func (e Sheet) SetProcess(Title string) core.Sheeter {
	e.Process = Title
	return e
}

// New creates new instances of Sheet
//
// Pre-cond: given process title, title of the Sheet and SheetContent instance
//
// Post-cond: return new instance of Sheet
func New(Process string, Title string) Sheet {
	return Sheet{
		Process: Process,
		Title:   Title,
		Content: newContent(),
	}
}

// Add adds new row to Sheet
//
// Pre-cond: given row instance and Sheet instance to add row
//
// Post-cond: row was added
func Add(row SheetRow, shet Sheet) Sheet {
	shet.Content = addRow(row, shet.Content)
	return shet
}
