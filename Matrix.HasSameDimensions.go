package matrix

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// HasSameDimensions - return whether the two matrices have the same dimensions.
//
// This will lock the underlying matrix
func (lhs *Matrix[T]) HasSameDimensions(rhs *Matrix[T]) error {
	lhs.lock.RLock()
	defer lhs.lock.RUnlock()

	rhs.lock.RLock()
	defer rhs.lock.RUnlock()

	return lhs.hasSameDimensions(rhs)
}

// hasSameDimensions - return whether the two matrices have the same dimensions.
//
//	This method is not exported because it is unsafe.  It has no locking
//	mechanism.
func (lhs *Matrix[T]) hasSameDimensions(rhs *Matrix[T]) error {

	if lhs == nil || rhs == nil || lhs.rows() != rhs.rows() || lhs.cols() != rhs.cols() {
		return fmt.Errorf(errors.MatrixDimensionMismatch)
	}
	return nil
}
