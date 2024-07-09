package matrix

// Add - Given matrix lhs, add matrix rhs and return the result matrix by reference.
//
//	This method will return if either input is nil or the two inputs are not of the same size.
//	We never want to return bad data that could be used.  This is the reason we return nil when
//	an error state has occurred.  This prevents the user from swallowing or otherwise ignoring
//	the error and some poor soul downstream from getting inaccurate information.  At least in
//	this case, the application will fail-safe.
func (lhs *Matrix[T]) Add(rhs *Matrix[T]) (result *Matrix[T], err error) {

	lhs.lock.RLock()
	defer lhs.lock.RUnlock()

	rhs.lock.RLock()
	defer rhs.lock.RUnlock()

	if err = lhs.hasSameDimensions(rhs); err != nil {
		return nil, err
	}

	if result, err = NewMatrix[T](lhs.rows(), lhs.cols()); err != nil {
		return nil, err
	}

	for r := uint(0); r < lhs.rows(); r++ {

		for c := uint(0); c < rhs.cols(); c++ {

			(*result).data[c][r] = (*lhs).data[c][r] + (*rhs).data[c][r]

		}

	}

	return result, nil

}
