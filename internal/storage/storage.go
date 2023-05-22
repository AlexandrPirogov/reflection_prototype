package storage

import (
	"reflection_prototype/internal/core/auth/user"
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/quant"
	"reflection_prototype/internal/core/sheet"
	"reflection_prototype/internal/core/thread"
)

type Storer interface {
	Login(u user.User) (string, error)
	Register(u user.User) error

	ReadProcess(p process.Process) (process.Process, error)
	ListProcesses() ([]process.Process, error)
	StoreProcess(p process.Process) error

	ReadThread(t thread.Thread) (thread.Thread, error)
	ListThreads() ([]thread.Thread, error)
	StoreThread(t thread.Thread) error

	ReadQuant(q quant.Quant) (quant.Quant, error)
	ListQuants() ([]quant.Quant, error)
	StoreQuant(q quant.Quant) error

	ListProcessesThreads(p process.Process) ([]thread.Thread, error)

	StoreSheet(s sheet.Sheet, p process.Process) error
	ReadSheet(process.Process) (sheet.Sheet, error)
	AddRow(r sheet.SheetRow, s sheet.Sheet) error
}
