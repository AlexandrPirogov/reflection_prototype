package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncorrectTitle(t *testing.T) {
	cases := []string{" startsspace", "with spaces", "123startsnum", "Qwe12!@#"}

	for _, v := range cases {
		t.Run(v, func(t *testing.T) {
			err := ValidateTitle(v)

			assert.NotNil(t, err)
		})
	}
}
