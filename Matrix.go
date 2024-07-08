package matrix

import (
	"github.com/sam-caldwell/types"
	"sync"
)

// Matrix - a generic two-dimensional float64 matrix
//
//		A matrix consists of a 2-dimensional field of elements of some generic type (types.Number)
//		which includes any built-in numeric type (e.g. int, int8,...uint, uint8, ...float32, float64).
//
//	 This object supports basic matrix arithmetic, and for certain types (integer types) bitwise
//	 operations are also supported.
type Matrix[MatrixElements types.Number] struct {
	lock sync.RWMutex

	data [][]MatrixElements
}
