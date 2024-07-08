//go:build !cuda
// +build !cuda

package matrix

import (
	"sync"
)

// Matrix - represents a 2D floating-point matrix (no CUDA)
type Matrix struct {
	lock sync.RWMutex
	data Simple2dMatrix[float64]
}
