package process

import (
	"fmt"
	"reflection_prototype/internal/core/thread"
	"reflection_prototype/internal/core/validator"
)

type Process struct {
	title   string
	threads map[string]thread.Thread
}

// New creates new instance of Process with given name
//
// Pre-cond: given title
//
// Post-cond: created new instance of Process
func New(title string) (Process, error) {
	if !validator.ValidateTitle(title) {
		return Process{}, fmt.Errorf("not valid title given")
	}
	return Process{
		title:   title,
		threads: make(map[string]thread.Thread),
	}, nil
}

// Title returns title of given Process instance
//
// Pre-cond: given instance of Process
//
// Post-cond: return title of Process
func Title(p Process) string {
	return p.title
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
