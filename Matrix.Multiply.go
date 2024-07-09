package matrix

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Multiply - multiply lhs into rhs and store the result by reference.
//
//	This supports any int, uint or float type.
//	Result cannot be nil or a crash will occur.
func (lhs *Matrix[T]) Multiply(rhs *Matrix[T]) (result *Matrix[T], err error) {

	lhs.lock.RLock()
	defer lhs.lock.RUnlock()

	rhs.lock.RLock()
	defer rhs.lock.RUnlock()

	if lhs.cols() != rhs.rows() {
		return nil, fmt.Errorf(errors.MatrixDimensionMismatch)
	}

	if result, err = NewMatrix[T](lhs.rows(), lhs.cols()); err != nil {
		return nil, err
	}

	for r := uint(0); r < lhs.rows(); r++ {

		for c := uint(0); c < rhs.cols(); c++ {

			var sum T = 0

			for lc := uint(0); lc < lhs.cols(); lc++ {

				sum += (*lhs).data[r][lc] * (*rhs).data[lc][c]

			}

			(*result).data[c][r] = sum

		}

	}
	return result, nil
}
