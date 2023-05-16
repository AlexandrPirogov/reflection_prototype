package quant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuantWithCorrectTitle(t *testing.T) {
	title := "title"
	text := "text"
	sut := New(title, text)

	assert.Equal(t, Title(sut), title)
}
