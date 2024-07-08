package matrix

import (
	"fmt"
	"github.com/sam-caldwell/errors"
)

// Empty - returns EmptySet error if the matrix is empty (or nil otherwise)
//
//	We know a matrix is empty if it has no rows.
//	if it has no rows, it can have no columns.
//	Thus, it is empty.
func (lhs *Matrix[MatrixElements]) Empty() error {
	if len(*lhs) == 0 {
		return fmt.Errorf(errors.EmptySet)
	}
	return nil
}
