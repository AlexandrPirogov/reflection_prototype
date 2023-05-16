package thread

import (
	"fmt"
	"reflection_prototype/internal/core/quant"
)

type Thread struct {
	title  string
	quants map[string]quant.Quant
}

// New creates new instance of thread with given name
//
// Pre-cond: given title
//
// Post-cond: created new instance of Thread
func New(title string) (Thread, error) {
	return Thread{
		title:  title,
		quants: make(map[string]quant.Quant),
	}, nil
}

// Title returns title of given Thread instance
//
// Pre-cond: given instance of Thread
//
// Post-cond: return title of Thread
func Title(t Thread) string {
	return t.title
}

// Add adds new Quant to Thread
//
// Pre-cond: given quant to add and thread in which to add quant
//
// Post-cond: if given quant isn't exists in Thread -- quant was added
// if given quant exists -- returns ??? #TODO
func Add(q quant.Quant, t Thread) (Thread, error) {
	_, ok := Seek(q, t)
	if ok {
		return t, fmt.Errorf("quant already exists with given name")
	}

	quantTitle := quant.Title(q)
	t.quants[quantTitle] = q
	return t, nil
}

// Seek seeks for given quant in given thread
//
// Pre-cond: given quant and thread in which to seek quant
//
// Post-cond: if quant exists -- returns quant and true
// Otherwise returns Quant{} and false
func Seek(q quant.Quant, t Thread) (quant.Quant, bool) {
	quantTitle := quant.Title(q)
	if v, ok := t.quants[quantTitle]; ok {
		return v, ok
	}

	return quant.Quant{}, false
}

// Len returns count of stored quants
//
// Pre-cond: given thread
//
// Post-cond: returned count of stored quants
func Len(t Thread) int {
	return len(t.quants)
}
