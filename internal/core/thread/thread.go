package thread

import "reflection_prototype/internal/core/quant"

type Thread struct {
	Quants map[string]quant.Quant
}
