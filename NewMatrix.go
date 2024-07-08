package matrix

import (
	"github.com/sam-caldwell/types"
)

// NewMatrix - Generate a new generic matrix of row x col using type T.
func NewMatrix[T types.Number](rows, cols uint) *Matrix[T] {

	var m Matrix[T]

	m.data = make([][]T, rows)

	for row := uint(0); row < m.rows(); row++ {

		m.data[row] = make([]T, cols)

	}

	return &m

}
