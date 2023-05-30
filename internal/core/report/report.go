package report

import (
	"reflection_prototype/internal/core/sheet"
)

type Report struct {
	Title   string `json:"Title"`
	Content map[string]string
}

// New creates new instance of Report
//
// Pre-cond:
//
// Post-cond: new instance of Report was created
func New(Title string) Report {
	return Report{
		Title:   Title,
		Content: map[string]string{},
	}
}

// Add adds new sheetrow to given Report
//
// Pre-cond: given SheetRow to add and report to which add
//
// Post-cond: row was added to report. If Row already exists in
// Report then returns same report
func Add(s sheet.SheetRow, r Report) Report {
	if r.Content == nil {
		r.Content = map[string]string{}
	}
	if _, ok := r.Content[s.Theme]; !ok {
		r.Content[s.Theme] = s.Spent
		return r
	}

	return r
}

// Remove removes given row from given report
//
// Pre-cond: given row and report to remove row from
//
// Post-cond: row was removed if it existed in report
// Otherwise returns given report
func Remove(s sheet.SheetRow, r Report) Report {
	if _, ok := r.Content[s.Theme]; ok {
		delete(r.Content, s.Theme)
		return r
	}

	return r
}
