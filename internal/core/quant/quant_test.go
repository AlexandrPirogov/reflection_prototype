package quant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuantWithCorrectTitle(t *testing.T) {
	title := "title"

	sut := New(title)

	assert.Equal(t, Title(sut), title)
}
