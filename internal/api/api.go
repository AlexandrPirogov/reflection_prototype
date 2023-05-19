package api

import "reflection_prototype/internal/core/process"

type Storer interface {
	StoreProcess(p process.Process) error
	ReadProcesses(p process.Process) ([]process.Process, error)
}

type Handler struct {
	S Storer
}
