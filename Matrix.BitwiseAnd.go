package matrix

import "github.com/sam-caldwell/errors"

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

	for r := uint(0); r < lhs.rows(); r++ {

		for c := uint(0); c < rhs.cols(); c++ {
			switch v := any((*lhs).data[r][c]).(type) {
			case int:
				(*result).data[r][c] = any(v & any((*rhs).data[r][c]).(int)).(T)
			case int8:
				(*result).data[r][c] = any(v & any((*rhs).data[r][c]).(int8)).(T)
			case int16:
				(*result).data[r][c] = any(v & any((*rhs).data[r][c]).(int16)).(T)
			case int32:
				(*result).data[r][c] = any(v & any((*rhs).data[r][c]).(int32)).(T)
			case int64:
				(*result).data[r][c] = any(v & any((*rhs).data[r][c]).(int64)).(T)
			case uint:
				(*result).data[r][c] = any(v & any((*rhs).data[r][c]).(uint)).(T)
			case uint8:
				(*result).data[r][c] = any(v & any((*rhs).data[r][c]).(uint8)).(T)
			case uint16:
				(*result).data[r][c] = any(v & any((*rhs).data[r][c]).(uint16)).(T)
			case uint32:
				(*result).data[r][c] = any(v & any((*rhs).data[r][c]).(uint32)).(T)
			case uint64:
				(*result).data[r][c] = any(v & any((*rhs).data[r][c]).(uint64)).(T)
			default:
				panic(errors.UnsupportedType)
			}
		}
	}
	return result, nil
}
