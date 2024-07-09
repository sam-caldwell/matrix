package matrix

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// And - Perform a bitwise AND operation on two Matrix objects and return the result
//
//	This supports any int, uint type.
func (lhs *Matrix[T]) And(rhs *Matrix[T]) (result *Matrix[T], err error) {

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
			switch v := any((*lhs).data[c][r]).(type) {
			case int:
				(*result).data[c][r] = any(v & any((*rhs).data[c][r]).(int)).(T)
			case int8:
				(*result).data[c][r] = any(v & any((*rhs).data[c][r]).(int8)).(T)
			case int16:
				(*result).data[c][r] = any(v & any((*rhs).data[c][r]).(int16)).(T)
			case int32:
				(*result).data[c][r] = any(v & any((*rhs).data[c][r]).(int32)).(T)
			case int64:
				(*result).data[c][r] = any(v & any((*rhs).data[c][r]).(int64)).(T)
			case uint:
				(*result).data[c][r] = any(v & any((*rhs).data[c][r]).(uint)).(T)
			case uint8:
				(*result).data[c][r] = any(v & any((*rhs).data[c][r]).(uint8)).(T)
			case uint16:
				(*result).data[c][r] = any(v & any((*rhs).data[c][r]).(uint16)).(T)
			case uint32:
				(*result).data[c][r] = any(v & any((*rhs).data[c][r]).(uint32)).(T)
			case uint64:
				(*result).data[c][r] = any(v & any((*rhs).data[c][r]).(uint64)).(T)
			default:
				return nil, fmt.Errorf(errors.UnsupportedType)
			}
		}
	}
	return result, nil
}
