//go:build !cuda
// +build !cuda

package matrix

// NewMatrix creates a new matrix with the given dimensions.
func NewMatrix(rows, cols int) *Matrix {

	return &Matrix{
		data: NewSimple2d(rows, cols),
	}

}
