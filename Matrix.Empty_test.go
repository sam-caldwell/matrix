package matrix

import (
	"github.com/sam-caldwell/errors"
	"testing"
)

func TestMatrix_Empty(t *testing.T) {

	t.Run("happy path: create non-empty matrix", func(t *testing.T) {
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

		if err := m.Empty(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("sad path: create empty matrix", func(t *testing.T) {
		const (
			rowSize = 5
			colSize = 2
		)
		m := Matrix[int]{}

		if err := m.Empty(); err == nil {
			t.Fatal("expected empty set error")
		} else {
			if err.Error() != errors.EmptySet {
				t.Fatalf("unexpected error: %v", err)
			}
		}
	})

}
