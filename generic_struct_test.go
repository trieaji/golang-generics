package golanggenerics

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T) {
	data := Data[string]{
		First: "Laksa",
		Second: "Giyuu",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "Laksa",
		Second: "Bayu",
	}

	assert.Equal(t, "Budi", data.ChangeFirst("Budi"))
	assert.Equal(t, "Hello Laksa", data.SayHello("Laksa"))
}