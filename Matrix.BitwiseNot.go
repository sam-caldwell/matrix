package matrix

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Not - Perform a bitwise NOT operation on a Matrix
//
//	This supports any int, uint type.
func (lhs *Matrix[T]) Not() (result *Matrix[T], err error) {

	lhs.lock.RLock()
	defer lhs.lock.RUnlock()

	if result, err = NewMatrix[T](lhs.rows(), lhs.cols()); err != nil {
		return nil, err
	}

	for r := uint(0); r < lhs.rows(); r++ {

		for c := uint(0); c < lhs.cols(); c++ {
			switch v := any((*lhs).data[c][r]).(type) {
			case int:
				result.data[c][r] = any(^v).(T)
			case int8:
				result.data[c][r] = any(^v).(T)
			case int16:
				result.data[c][r] = any(^v).(T)
			case int32:
				result.data[c][r] = any(^v).(T)
			case int64:
				result.data[c][r] = any(^v).(T)
			case uint:
				result.data[c][r] = any(^v).(T)
			case uint8:
				result.data[c][r] = any(^v).(T)
			case uint16:
				result.data[c][r] = any(^v).(T)
			case uint32:
				result.data[c][r] = any(^v).(T)
			case uint64:
				result.data[c][r] = any(^v).(T)
			default:
				return nil, fmt.Errorf(errors.UnsupportedType)
			}
		}
	}
	return result, nil
}
