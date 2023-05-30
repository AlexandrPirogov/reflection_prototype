package thread

import (
	"reflection_prototype/internal/core/quant"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewThreadWithCorrectTitle(t *testing.T) {
	title := "some_title"
	process := "some_process"

	sut, _ := New(process, title)

	assert.Equal(t, Title(sut), title)
}

func TestAddNotExistingQuantsToThread(t *testing.T) {
	q1, _ := quant.New("proc", "thread", "somebody", "once")
	q2, _ := quant.New("proc", "thread", "told", "me")
	q3, _ := quant.New("proc", "thread", "the", "world")
	cases := []quant.Quant{q1, q2, q3}

	sut, _ := New("s", "sut")
	var err error

	for _, q := range cases {
		t.Run(quant.Title(q), func(t *testing.T) {
			sut, err = Add(q, sut)

			assert.Nil(t, err)
		})
	}
}

func TestAddWithExistingQuantsToThread(t *testing.T) {
	q1, _ := quant.New("proc", "thread", "somebody", "once")
	q2, _ := quant.New("proc", "thread", "told", "me")
	q3, _ := quant.New("proc", "thread", "the", "world")
	cases := []quant.Quant{q1, q2, q3}

	sut, _ := New("s", "sut")
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
