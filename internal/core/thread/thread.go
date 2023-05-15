package thread

import "reflection_prototype/internal/core/quant"

type Thread struct {
	title  string
	quants map[string]quant.Quant
}

// New creates new instance of thread with given name
//
// Pre-cond: given title
//
// Post-cond: created new instance of Thread
func New(title string) Thread {
	return Thread{
		title:  title,
		quants: make(map[string]quant.Quant),
	}
}

// Title returns title of given Thread instance
//
// Pre-cond: given instance of Thread
//
// Post-cond: return title of Thread
func Title(t Thread) string {
	return t.title
}
