package matrix

import (
	"github.com/sam-caldwell/errors"
	"testing"
)

func TestMatrix_Get(t *testing.T) {

	createEmptyMatrix := func() *Matrix[int] {
		return &(Matrix[int]{})
	}

	fillMatrix := func(m *Matrix[int]) {
		i := 0
		for r := uint(0); r < m.rows(); r++ {
			for c := uint(0); c < m.cols(); c++ {
				(*m).data[r][c] = i
				i++
			}
		}
	}

	t.Run("Happy Path: Given a matrix 5x5, with unique data 0-25, data reads ok", func(t *testing.T) {
		var m *Matrix[int]
		t.Run("Setup the matrix", func(t *testing.T) {
			var err error
			if m, err = NewMatrix[int](5, 5); err != nil {
				t.Fatal(err)
			}
			if m == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrix(m)
		})
		t.Run("Verify the matrix", func(t *testing.T) {
			for r := uint(0); r < m.rows(); r++ {
				for c := uint(0); c < m.cols(); c++ {
					expected := int(c*m.rows() + r)
					actual, err := m.Get(r, c)
					if err != nil {
						t.Fatalf("unexpected error: %v at (%d,%d) with size (%d,%d)",
							err, r, c, m.rows(), m.cols())
					}
					if actual != expected {
						t.Fatalf("mismatched values\n"+
							"expected: %v\nactual  : %v",
							expected, actual)
					}
				}
			}
		})
	})

	t.Run("Sad Path: Given an empty matrix, we will get an BoundsCheckError error", func(t *testing.T) {
		var m = createEmptyMatrix()
		if m == nil {
			t.Fatal("unexpected nil matrix")
		}
		if v, err := m.Get(0, 0); err == nil {
			t.Fatal("expected error.  none found")
		} else {
			if err.Error() != errors.BoundsCheckError {
				t.Fatalf("expected BoundsCheckError but got %v", err)
			}
			if v != 0 {
				t.Fatal("expect zero-value on error")
			}
		}
	})

	t.Run("Sad Path: given 5x5 matrix, reading out of bounds will return BoundsCheckError", func(t *testing.T) {
		var m *Matrix[int]
		t.Run("Setup the matrix", func(t *testing.T) {
			var err error
			if m, err = NewMatrix[int](5, 5); err != nil {
				t.Fatal(err)
			}
			if m == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrix(m)
		})
		t.Run("Read out of bounds coordinates and verify we get a BoundsCheckError", func(t *testing.T) {
			if v, err := m.Get(10, 10); err == nil {
				t.Fatal("expected error not found")
			} else {
				if err.Error() != errors.BoundsCheckError {
					t.Fatalf("expected BoundsCheckError. Got %v", err)
				}
				if v != 0 {
					t.Fatal("expect 0 if BoundsCheckError happens")
				}
			}
		})
	})
}
