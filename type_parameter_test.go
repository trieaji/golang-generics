package golanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T { //bisa seperti ini "[T interface{}]", "[T interface{}]" adalah type parameter
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	result := Length[string]("Law")
	fmt.Println(result)
	assert.Equal(t, "Law", result)

	resultNumber := Length[int](100)
	fmt.Println(resultNumber)
	assert.Equal(t, 100, resultNumber)
}

