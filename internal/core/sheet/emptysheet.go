package sheet

import "reflection_prototype/internal/core"

type EmptySheet struct {
	Process string
	Title   string
}

func (e EmptySheet) IsEmpty() bool {
	return true
}

func (e EmptySheet) SetTitle(Title string) core.Sheeter {
	e.Title = Title
	return e
}

func (e EmptySheet) SetProcess(Title string) core.Sheeter {
	e.Process = Title
	return e
}
