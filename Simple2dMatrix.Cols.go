package matrix

// Cols - return the number of columns
func (lhs *Simple2dMatrix[T]) Cols() int {
	if lhs.Rows() == 0 {
		return 0
	}
	return len((*lhs)[0])
}
