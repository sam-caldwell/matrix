package matrix

// Empty - returns boolean true if Matrix is empty.
//
//	We only check to see if there are any rows.  If there are any rows
//	it doesn't matter if there are any columns.
func (lhs *Simple2dMatrix[T]) Empty() bool {
	return uint(len(*lhs)) == 0
}
