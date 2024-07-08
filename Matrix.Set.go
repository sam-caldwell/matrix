package matrix

// Set - set the value at (row, col)
func (lhs *Matrix[T]) Set(row, col uint, value T) {

	lhs.lock.RLock()
	defer lhs.lock.RUnlock()

	(*lhs).data[row][col] = value

}
