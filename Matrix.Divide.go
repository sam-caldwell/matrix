package matrix

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Divide - divide lhs into rhs and return the result by reference.
//
//	   There are two error scenarios:
//		- either lhs or rhs is nil or the dimensions do not match
//	    - any divisor is zero (0) and division by zero is detected.
func (lhs *Matrix[T]) Divide(rhs *Matrix[T]) (result *Matrix[T], err error) {

	lhs.lock.RLock()
	defer lhs.lock.RUnlock()

	rhs.lock.RLock()
	defer rhs.lock.RUnlock()

	if err = lhs.hasSameDimensions(rhs); err != nil {
		return nil, err
	}

	if result, err = NewMatrix[T](lhs.rows(), lhs.cols()); err != nil {
		return nil, err
	}

	for r := uint(0); r < lhs.rows(); r++ {

		for c := uint(0); c < rhs.cols(); c++ {

			if (*rhs).data[c][r] == 0 {
				return nil, fmt.Errorf(errors.DivisionByZero)
			}

			(*result).data[c][r] = (*lhs).data[c][r] / (*rhs).data[c][r]

		}

	}

	return result, nil

}
