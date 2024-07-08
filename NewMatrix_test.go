package matrix

import (
	"github.com/sam-caldwell/errors"
	"testing"
)

func Test_NewMatrix(t *testing.T) {
	t.Run("happy path: simple creation of a matrix", func(t *testing.T) {
		const (
			rowSize = 5
			colSize = 2
		)

		m, err := NewMatrix[int](rowSize, colSize)
		if err != nil {
			t.Fatal(err)
		}

		if actual := len(m.data); actual != rowSize {
			t.Fatalf("rows mismatch.  got: %d", actual)
		}

		if actual := len(m.data[0]); actual != colSize {
			t.Fatalf("col mismatch.  got: %d", actual)
		}
	})
	t.Run("sad path: zero-size matrix should throw error and return nil Matrix pointer", func(t *testing.T) {
		const (
			rowSize = 0
			colSize = 0
		)
		m, err := NewMatrix[int](rowSize, colSize)
		if err == nil {
			t.Fatal("expected error")
		}
		if err.Error() != errors.BoundsCheckError {
			t.Fatalf("expected bounds check error")
		}
		if m != nil {
			t.Fatalf("expect nil pointer on error")
		}
	})
}
