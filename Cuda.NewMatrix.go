//go:build cuda
// +build cuda

package matrix

import "gocv.io/x/gocv"

// NewMatrix creates a new matrix with the given dimensions.
func NewMatrix(rows, cols int) *Matrix {
	data := gocv.NewMatWithSize(rows, cols, gocv.MatTypeCV64F)
	return &Matrix{
		data: data,
	}
}
