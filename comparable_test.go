package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//comparable adalah type paramater yg bisa digunakan sebagai perbandingan, kalau type parameternya menggunakan any tidak bisa

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[string]("Laksa", "Laksa"))
	assert.True(t, IsSame[int](100, 100))
	assert.True(t, IsSame[bool](false, false))
}