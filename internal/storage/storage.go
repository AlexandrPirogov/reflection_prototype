package storage

import (
	"reflection_prototype/internal/core/auth/user"
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/quant"
	"reflection_prototype/internal/core/report"
	"reflection_prototype/internal/core/sheet"
	"reflection_prototype/internal/core/thread"
)

type Storer interface {
	Login(u user.User) (string, error)
	Register(u user.User) error

	ReadProcess(u user.User, p process.Process) (process.Process, error)
	ListProcesses(u user.User) ([]process.Process, error)
	StoreProcess(u user.User, p process.Process) error

	ReadThread(u user.User, t thread.Thread) (thread.Thread, error)
	ListThreads(u user.User) ([]thread.Thread, error)
	StoreThread(u user.User, t thread.Thread) error

	ReadQuant(u user.User, q quant.Quant) (quant.Quant, error)
	ListQuants(u user.User) ([]quant.Quant, error)
	StoreQuant(u user.User, q quant.Quant) error

	ListProcessesThreads(u user.User, p process.Process) ([]thread.Thread, error)

	StoreSheet(u user.User, s sheet.Sheet, p process.Process) error
	ReadSheet(u user.User, p process.Process) (sheet.Sheet, error)
	AddRow(u user.User, r sheet.SheetRow, p process.Process) error
	MarkRow(u user.User, r sheet.SheetRow, p process.Process) error

	StartWork(u user.User, r sheet.SheetRow, p process.Process) error
	StopWork(u user.User, r sheet.SheetRow, p process.Process) error

	CreateReport(u user.User, r report.Report) error
	FillReport(u user.User, s sheet.SheetRow, p process.Process, r report.Report) error
	ReadReport(u user.User, r report.Report) (report.Report, error)
	//EaseReport(u user.User, s sheet.SheetRow, r report.Report) error
}
