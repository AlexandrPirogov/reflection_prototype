package quant

type Quant struct {
	title string
	text  string
}

// New creates new instance of Quant with given name
//
// Pre-cond: given title
//
// Post-cond: created new instance of Quant
func New(title, text string) Quant {
	return Quant{
		title: title,
		text:  text,
	}
}

// Title returns title of given Quant instance
//
// Pre-cond: given instance of Thread
//
// Post-cond: return title of Quant
func Title(q Quant) string {
	return q.title
}
