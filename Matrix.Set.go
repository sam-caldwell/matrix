package matrix

// Set - set the value at (row, col)
func (lhs *Matrix[T]) Set(row, col uint, value T) error {

	lhs.lock.RLock()
	defer lhs.lock.RUnlock()

	if err := lhs.boundsCheck(row, col); err != nil {
		return err
	}

	(*lhs).data[col][row] = value

	return nil
}
