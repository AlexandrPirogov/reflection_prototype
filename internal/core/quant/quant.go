package quant

type Quant struct {
	title string
	text  string
}

func New(title string) Quant {
	return Quant{
		title: title,
	}
}

func Title(q Quant) string {
	return q.title
}
