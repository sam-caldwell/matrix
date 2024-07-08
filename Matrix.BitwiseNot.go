package matrix

import "github.com/sam-caldwell/errors"

// Not - Perform a bitwise NOT operation on a Matrix
//
//	This supports any int, uint type.
func (lhs *Matrix[T]) Not() (result *Matrix[T], err error) {

	lhs.lock.RLock()
	defer lhs.lock.RUnlock()

	for r := uint(0); r < lhs.rows(); r++ {

		for c := uint(0); c < lhs.cols(); c++ {
			switch v := any((*lhs).data[r][c]).(type) {
			case int:
				result.data[r][c] = any(^v).(T)
			case int8:
				result.data[r][c] = any(^v).(T)
			case int16:
				result.data[r][c] = any(^v).(T)
			case int32:
				result.data[r][c] = any(^v).(T)
			case int64:
				result.data[r][c] = any(^v).(T)
			case uint:
				result.data[r][c] = any(^v).(T)
			case uint8:
				result.data[r][c] = any(^v).(T)
			case uint16:
				result.data[r][c] = any(^v).(T)
			case uint32:
				result.data[r][c] = any(^v).(T)
			case uint64:
				result.data[r][c] = any(^v).(T)
			default:
				panic(errors.UnsupportedType)
			}
		}
	}
	return result, nil
}
