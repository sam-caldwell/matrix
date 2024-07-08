//go:build cuda
// +build cuda

package matrix

import (
	"gocv.io/x/gocv"
	"sync"
)

// Matrix - represents a 2D floating-point matrix (for CUDA)
type Matrix struct {
	lock sync.RWMutex
	data gocv.Mat
}
