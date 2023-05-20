package storage

import (
	"reflection_prototype/internal/core/auth/user"
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/quant"
	"reflection_prototype/internal/core/thread"
)

type Storer interface {
	Login(u user.User) (string, error)
	Register(u user.User) error

	StoreProcess(p process.Process) error
	ReadProcesses(p process.Process) ([]process.Process, error)

	StoreThread(t thread.Thread) error
	ReadThreads(t thread.Thread) ([]thread.Thread, error)

	StoreQuant(q quant.Quant) error
	ReadQuants(q quant.Quant) ([]quant.Quant, error)
}
