package quant

import (
	"fmt"
	"reflection_prototype/internal/core/validator"
)

type Quant struct {
	title string
	text  string
}

// New creates new instance of Quant with given name
//
// Pre-cond: given title
//
// Post-cond: created new instance of Quant
func New(title, text string) (Quant, error) {
	if !validator.ValidateTitle(title) {
		return Quant{}, fmt.Errorf("not valid title given")
	}
	return Quant{
		title: title,
		text:  text,
	}, nil
}

// Title returns title of given Quant instance
//
// Pre-cond: given instance of Thread
//
// Post-cond: return title of Quant
func Title(q Quant) string {
	return q.title
}
