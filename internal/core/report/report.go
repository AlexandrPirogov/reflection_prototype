package report

import (
	"time"
)

type Report struct {
	Content map[string]time.Duration
}

func New() Report {
	return Report{}
}
