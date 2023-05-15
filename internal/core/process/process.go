package process

import "reflection_prototype/internal/core/thread"

type Process struct {
	Thread map[string]thread.Thread
}
