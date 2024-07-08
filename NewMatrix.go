package matrix

import (
	"fmt"
	"github.com/sam-caldwell/errors"
	"github.com/sam-caldwell/types"
)

// NewMatrix - Generate a new generic matrix of (row X col) using type T.
func NewMatrix[T types.Number](rows, cols uint) (result *Matrix[T], err error) {

	var m Matrix[T]

	if rows == 0 || cols == 0 {
		return nil, fmt.Errorf(errors.BoundsCheckError)
	}

	m.data = make([][]T, rows)

	for row := uint(0); row < m.rows(); row++ {

		m.data[row] = make([]T, cols)

	}

	return &m, nil

}
