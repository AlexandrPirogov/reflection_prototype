package core

type CoreType int

const (
	PROCESS CoreType = iota
	THREAD
	SHEET
	SHEETROW
)

type Container interface {
	IsEmpty() bool
}

type Sheeter interface {
	SetTitle(Title string) Sheeter
	SetProcess(Title string) Sheeter
	IsEmpty() bool
}
