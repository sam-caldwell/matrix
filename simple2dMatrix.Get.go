package matrix

// Get - return the value at (row, col)
func (lhs *Simple2dMatrix[T]) Get(row, col uint) T {
	return (*lhs)[row][col]
}
