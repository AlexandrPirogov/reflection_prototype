package storage

import (
	"reflection_prototype/internal/core/process"
)

type Storer interface {
	StoreProcess(p process.Process) error
}
