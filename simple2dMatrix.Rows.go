package matrix

// Rows - return the number of rows.
func (lhs *Simple2dMatrix[T]) Rows() int {

	return len(*lhs)

}
