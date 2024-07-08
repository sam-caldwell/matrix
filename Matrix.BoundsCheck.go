package matrix

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// BoundsCheck - return boolean true if given (row,col) are within bounds of the lhs matrix
//
// This will lock the underlying matrix
func (lhs *Matrix[T]) BoundsCheck(row, col uint) error {

	lhs.lock.RLock()
	defer lhs.lock.RUnlock()

	return lhs.boundsCheck(row, col)

}

// boundsCheck - return boolean true if given (row,col) are within bounds of the lhs matrix
//
//	This method is not exported because it is unsafe.  It has no locking
//	mechanism.
func (lhs *Matrix[T]) boundsCheck(row, col uint) error {

	if row < lhs.rows() && col < lhs.cols() {
		return nil
	}

	return fmt.Errorf(errors.BoundsCheckError)

}
