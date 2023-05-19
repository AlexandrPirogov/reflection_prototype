package process

import (
	"fmt"
	"reflection_prototype/internal/core/thread"
	"reflection_prototype/internal/validator"
)

type Process struct {
	Title   string `json:"Title"`
	threads map[string]thread.Thread
}

// New creates new instance of Process with given name
//
// Pre-cond: given Title
//
// Post-cond: created new instance of Process
func New(Title string) (Process, error) {
	if !validator.ValidateTitle(Title) {
		return Process{}, fmt.Errorf("not valid Title given")
	}
	return Process{
		Title:   Title,
		threads: make(map[string]thread.Thread),
	}, nil
}

// Title returns Title of given Process instance
//
// Pre-cond: given instance of Process
//
// Post-cond: return Title of Process
func Title(p Process) string {
	return p.Title
}

// Seek seeks for given thread in given process
//
// Pre-cond: given thread and process in which to seek quant
//
// Post-cond: if thread exists -- returns thread and true
// Otherwise returns Thread{} and false
func Seek(t thread.Thread, p Process) (thread.Thread, bool) {
	threadTitle := thread.Title(t)
	if t, ok := p.threads[threadTitle]; ok {
		return t, ok
	}
	return thread.Thread{}, false
}

// Add adds new Thread to Process
//
// Pre-cond: given thread to add and process in which to add the thread
//
// Post-cond: if given thread isn't exists in Process -- thread was added; return thread and nil
// if given thread exists -- returns errors
func Add(t thread.Thread, p Process) (Process, error) {
	_, ok := Seek(t, p)
	if ok {
		return p, fmt.Errorf("quant already exists with given name")
	}

	threadTitle := thread.Title(t)
	p.threads[threadTitle] = t
	return p, nil
}
