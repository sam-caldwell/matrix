package matrix

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Divide - divide lhs into rhs and return the result by reference.
//
//	   There are two error scenarios:
//			- either lhs or rhs is nil or the dimensions do not match
//	     - any divisor is zero (0) and division by zero is detected.
func (lhs *Matrix[T]) Divide(rhs *Matrix[T]) (result *Matrix[T], err error) {

	lhs.lock.RLock()
	defer lhs.lock.RUnlock()

	rhs.lock.RLock()
	defer rhs.lock.RUnlock()

	if err = lhs.hasSameDimensions(rhs); err != nil {
		return nil, err
	}

	result = NewMatrix[T](lhs.rows(), lhs.cols())

	for r := uint(0); r < lhs.rows(); r++ {

		for c := uint(0); c < rhs.cols(); c++ {

			if (*rhs).data[r][c] == 0 {
				return nil, fmt.Errorf(errors.DivisionByZero)
			}

			(*result).data[r][c] = (*lhs).data[r][c] / (*rhs).data[r][c]

		}

	}

	return result, nil

}
