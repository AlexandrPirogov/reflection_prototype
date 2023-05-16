package process

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateProcessWithCorrectTitle(t *testing.T) {
	title := "some_title"

	sut, _ := New(title)

	assert.Equal(t, title, sut.title)
}
