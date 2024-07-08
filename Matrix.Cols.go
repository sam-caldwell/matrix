package matrix

// Cols - return the number of columns.
//
// This will lock the underlying matrix
func (lhs *Matrix[T]) Cols() uint {

	lhs.lock.RLock()
	defer lhs.lock.RUnlock()

	return lhs.cols()

}

// cols - return the number of columns
//
//	This method is not exported because it is unsafe.  It has no locking
//	mechanism.
func (lhs *Matrix[T]) cols() uint {

	if lhs.rows() == 0 {
		return 0
	}

	return uint(len((*lhs).data[0]))

}
