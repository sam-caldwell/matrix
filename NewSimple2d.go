package matrix

import (
	"github.com/sam-caldwell/types"
)

// NewSimple2d - Generate a new simple matrix of row x col using type T.
func NewSimple2d[T types.Number](rows, cols int) *Simple2dMatrix[T] {
	var m Simple2dMatrix[T]
	m = make(Simple2dMatrix[T], rows)
	for i, _ := range m {
		m[i] = make([]T, cols)
	}
	return &m
}
