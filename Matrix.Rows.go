package matrix

// Rows - return the number of rows.
//
// This will lock the underlying matrix
func (lhs *Matrix[T]) Rows() uint {

	lhs.lock.RLock()
	defer lhs.lock.RUnlock()

	return lhs.rows()

}

// rows - return the number of rows.
//
//	This method is not exported because it is unsafe.  It has no locking
//	mechanism.
func (lhs *Matrix[T]) rows() uint {

	return uint(len((*lhs).data))

}
