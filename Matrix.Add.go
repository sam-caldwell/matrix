package matrix

// Add - Given matrix lhs, add matrix rhs and return the result matrix by reference.
//
//	This method will return if either input is nil or the two inputs are not of the same size.
func (lhs *Matrix[T]) Add(rhs *Matrix[T]) (result *Matrix[T], err error) {

	lhs.lock.RLock()
	defer lhs.lock.RUnlock()

	rhs.lock.RLock()
	defer rhs.lock.RUnlock()

	if err = lhs.hasSameDimensions(rhs); err != nil {
		return nil, err
	}

	result = NewMatrix[T](lhs.rows(), lhs.cols())

	for r := uint(0); r < lhs.rows(); r++ {

		for c := uint(0); c < rhs.cols(); c++ {

			(*result).data[r][c] = (*lhs).data[r][c] + (*rhs).data[r][c]

		}

	}

	return result, nil

}
