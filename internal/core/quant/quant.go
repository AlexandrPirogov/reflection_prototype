package quant

import (
	"fmt"
	"reflection_prototype/internal/validator"
	"time"
)

type Quant struct {
	Process   string    `json:"process"`
	Thread    string    `json:"thread"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

// New creates new instance of Quant with given name
//
// Pre-cond: given Title
//
// Post-cond: created new instance of Quant
func New(Title, Text string) (Quant, error) {
	if !validator.ValidateTitle(Title) {
		return Quant{}, fmt.Errorf("not valid Title given")
	}
	return Quant{
		Title: Title,
		Text:  Text,
	}, nil
}

// Title returns Title of given Quant instance
//
// Pre-cond: given instance of Thread
//
// Post-cond: return Title of Quant
func Title(q Quant) string {
	return q.Title
}
