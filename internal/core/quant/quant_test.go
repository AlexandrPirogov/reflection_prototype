package quant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuantWithCorrectTitle(t *testing.T) {
	title := "title"
	text := "text"
	sut, _ := New(title, text)

	assert.Equal(t, Title(sut), title)
}

func TestNewQuantWithIncorrectTitle(t *testing.T) {
	cases := []string{" startsspace", "with spaces", "123startsnum", "Qwe12!@#"}

	for _, v := range cases {
		t.Run(v, func(t *testing.T) {
			_, err := New(v, "")

			assert.NotNil(t, err)
		})
	}
}
