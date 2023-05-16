package postgres

import (
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/quant"
	"reflection_prototype/internal/core/thread"
)

type PgConnection struct {
}

func (pg *PgConnection) StoreProcess(p process.Process) error {
	return nil
}

func (pg *PgConnection) SelectProcess(p process.Process) (process.Process, error) {
	return p, nil
}

func (pg *PgConnection) StoreThread(t thread.Thread, p process.Process) error {
	return nil
}
func (pg *PgConnection) SelectThread(t thread.Thread, p process.Process) (thread.Thread, error) {
	return t, nil
}

func (pg *PgConnection) StoreQuant(q quant.Quant, t thread.Thread) error {
	return nil
}
