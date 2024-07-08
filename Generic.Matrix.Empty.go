package matrix

// Empty - returns boolean true if Matrix is empty
func (lhs *Matrix) Empty() bool {
	return lhs.data.Empty()
}
