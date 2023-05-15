package process

import "reflection_prototype/internal/core/thread"

type Process struct {
	title   string
	threads map[string]thread.Thread
}

// New creates new instance of Process with given name
//
// Pre-cond: given title
//
// Post-cond: created new instance of Process
func New(title string) Process {
	return Process{
		title:   title,
		threads: make(map[string]thread.Thread),
	}
}

// Title returns title of given Process instance
//
// Pre-cond: given instance of Process
//
// Post-cond: return title of Process
func Title(p Process) string {
	return p.title
}
