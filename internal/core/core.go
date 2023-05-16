package core

import (
	"reflection_prototype/internal/core/process"
	"reflection_prototype/internal/core/quant"
	"reflection_prototype/internal/core/thread"
)

type Storage interface {
	StoreProcess(p process.Process) error
	SelectProcess(p process.Process) error

	StoreThread(t thread.Thread, p process.Process) error
	SelectThread(t thread.Thread) error

	StoreQuant(q quant.Quant, t thread.Thread) error
}

// CreateProcess creates new process with given name and stores it to the given Storage
//
// Pre-cond: given title for new process. Title is unique
//
// Post-cond: if title is unique -- process was saved in storage, nil returned; otherwise returns err
func CreateProcess(title string, s Storage) error {
	p, err := process.New(title)
	if err != nil {
		return err
	}

	err = s.StoreProcess(p)
	if err != nil {
		return err
	}

	return nil
}

// CreateThread creates new thread with given name and stores it to the given Storage
//
// Pre-cond: given process title, thread title and storage.
// Process must exists and title of thread must be unique
// Post-cond: thread was created, error returned nil. Otherwise returns err
func CreateThread(procTitle string, threadTitle string, s Storage) error {
	proc, err := process.New(procTitle)
	if err != nil {
		return err
	}

	t := thread.New(threadTitle)
	p, err := process.Add(t, proc)
	if err != nil {
		return err
	}

	err = s.StoreThread(t, p)
	if err != nil {
		return err
	}

	return nil
}
