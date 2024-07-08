package matrix

// HasSameDimensions - return whether the two matrices have the same dimensions.
func (lhs *Matrix) HasSameDimensions(rhs *Matrix) bool {
	return lhs.data.Rows() == rhs.data.Rows() || lhs.data.Cols() == rhs.data.Cols()
}
