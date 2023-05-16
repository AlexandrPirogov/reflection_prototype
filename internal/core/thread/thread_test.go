package thread

import (
	"reflection_prototype/internal/core/quant"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewThreadWithCorrectTitle(t *testing.T) {
	title := "some_title"

	sut, _ := New(title)

	assert.Equal(t, Title(sut), title)
}

func TestAddNotExistingQuantsToThread(t *testing.T) {
	cases := []quant.Quant{
		quant.New("somebody", "once"),
		quant.New("told", "me"),
		quant.New("the", "world"),
	}

	sut, _ := New("sut")
	var err error

	for _, q := range cases {
		t.Run(quant.Title(q), func(t *testing.T) {
			sut, err = Add(q, sut)

			assert.Nil(t, err)
		})
	}
}

func TestAddWithExistingQuantsToThread(t *testing.T) {
	cases := []quant.Quant{
		quant.New("somebody", "once"),
		quant.New("told", "me"),
		quant.New("the", "world"),
	}

	sut, _ := New("sut")
	var err error

	for _, q := range cases {
		sut, err = Add(q, sut)
	}

	for _, q := range cases {
		t.Run(quant.Title(q), func(t *testing.T) {
			sut, err = Add(q, sut)

			assert.NotNil(t, err)
		})
	}
}
