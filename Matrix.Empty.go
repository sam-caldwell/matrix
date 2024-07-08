package matrix

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Empty - returns EmptySet error if the matrix is empty (or nil otherwise)
//
//	We know a matrix is empty if it has no columns.
//	if it has no columns, it can have no rows.
//	Thus, it is empty.
func (lhs *Matrix[MatrixElements]) Empty() error {
	if len((*lhs).data) == 0 {
		return fmt.Errorf(errors.EmptySet)
	}
	return nil
}
