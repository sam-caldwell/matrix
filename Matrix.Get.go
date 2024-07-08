package matrix

// Get - return the value at (row, col).
func (lhs *Matrix[T]) Get(row, col uint) (T, error) {

	lhs.lock.RLock()
	defer lhs.lock.RUnlock()

	if err := lhs.boundsCheck(row, col); err != nil {
		return 0, err
	}

	return (*lhs).data[row][col], nil

}
