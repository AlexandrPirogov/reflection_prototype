package thread

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewThreadWithCorrectTitle(t *testing.T) {
	title := "some_title"

	sut := New(title)

	assert.Equal(t, Title(sut), title)
}
