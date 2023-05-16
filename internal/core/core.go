package core

import "reflection_prototype/internal/core/process"

// CreateProcess creates new process with given name
//
// Pre-cond: given title for new process. Title is unique
//
// Post-cond: if title is unique -- returns nil, otherwise returns err
func CreateProcess(title string) error {
	_, err := process.New(title)
	return err
}
