package matrix

import "github.com/sam-caldwell/types"

// Simple2dMatrix - a simple two-dimensional float64 matrix
type Simple2dMatrix[MatrixElements types.Number] [][]MatrixElements
