package matrix

import (
	"github.com/sam-caldwell/errors"
	"testing"
)

func TestMatrix_Not(t *testing.T) {

	fillMatrixInt := func(m *Matrix[int], v int) {
		for r := uint(0); r < m.rows(); r++ {
			for c := uint(0); c < m.cols(); c++ {
				(*m).data[r][c] = v
			}
		}
	}
	fillMatrixFloat32 := func(m *Matrix[float32], v float32) {
		i := v
		for r := uint(0); r < m.rows(); r++ {
			for c := uint(0); c < m.cols(); c++ {
				(*m).data[r][c] = v
				i++
			}
		}
	}
	fillMatrixFloat64 := func(m *Matrix[float64], v float64) {
		i := v
		for r := uint(0); r < m.rows(); r++ {
			for c := uint(0); c < m.cols(); c++ {
				(*m).data[r][c] = v
				i++
			}
		}
	}
	t.Run("Given 0's integer matrix A, perform C=Not(A), expect 1's without error", func(t *testing.T) {
		const (
			rowSize  = 5
			colSize  = 5
			valueA   = 0
			expected = ^valueA
		)
		var (
			A *Matrix[int]
			C *Matrix[int]
		)
		t.Run("Setup the matrix", func(t *testing.T) {
			var err error
			if A, err = NewMatrix[int](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if A == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixInt(A, valueA)
		})
		t.Run("C=Not(A)", func(t *testing.T) {
			var err error
			if C, err = A.Not(); err != nil {
				t.Fatal(err)
			}
			for r := uint(0); r < C.rows(); r++ {
				for c := uint(0); c < C.cols(); c++ {
					if actual := (*C).data[r][c]; actual != expected {
						t.Fatalf("value mismatch\n"+
							"actual:   %d\n"+
							"expected: %d", actual, expected)
					}
				}
			}
		})
	})

	t.Run("Given 1's integer matrix A, perform C=Not(A), expect 0's without error", func(t *testing.T) {
		const (
			rowSize  = 5
			colSize  = 5
			valueA   = 1
			expected = ^valueA
		)
		var (
			A *Matrix[int]
			C *Matrix[int]
		)
		t.Run("Setup the matrix", func(t *testing.T) {
			var err error
			if A, err = NewMatrix[int](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if A == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixInt(A, valueA)
		})
		t.Run("C=Not(A)", func(t *testing.T) {
			var err error
			if C, err = A.Not(); err != nil {
				t.Fatal(err)
			}
			for r := uint(0); r < C.rows(); r++ {
				for c := uint(0); c < C.cols(); c++ {
					if actual := (*C).data[r][c]; actual != expected {
						t.Fatalf("value mismatch\n"+
							"actual:   %d\n"+
							"expected: %d", actual, expected)
					}
				}
			}
		})
	})

	t.Run("Given 5's integer matrix A, perform C=Not(A), expect -6's without error", func(t *testing.T) {
		const (
			rowSize  = 5
			colSize  = 5
			valueA   = 5
			expected = ^valueA
		)
		var (
			A *Matrix[int]
			C *Matrix[int]
		)
		t.Run("Setup the matrix", func(t *testing.T) {
			var err error
			if A, err = NewMatrix[int](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if A == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixInt(A, valueA)
		})
		t.Run("C=Not(A).", func(t *testing.T) {
			var err error
			if C, err = A.Not(); err != nil {
				t.Fatal(err)
			}
			for r := uint(0); r < C.rows(); r++ {
				for c := uint(0); c < C.cols(); c++ {
					if actual := (*C).data[r][c]; actual != expected {
						t.Fatalf("value mismatch\n"+
							"actual:   %d\n"+
							"expected: %d", actual, expected)
					}
				}
			}
		})
	})

	t.Run("Given float32 matrix A, perform C=Not(A), expect UnsupportedType", func(t *testing.T) {
		const (
			rowSize = 5
			colSize = 5
			valueA  = float32(1.0)
		)
		var (
			A *Matrix[float32]
			C *Matrix[float32]
		)
		t.Run("Setup the matrix", func(t *testing.T) {
			var err error
			if A, err = NewMatrix[float32](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if A == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixFloat32(A, valueA)
		})
		t.Run("C=Not(A).  Expect Error", func(t *testing.T) {
			var err error
			if C, err = A.Not(); err == nil {
				t.Fatal("expected UnsupportedType error.  got none.")
			} else {
				if err.Error() != errors.UnsupportedType {
					t.Fatalf("expected UnsupportedType.  Got %v", err)
				}
				if C != nil {
					t.Fatalf("expected nil on error")
				}
			}
		})
	})

	t.Run("Given float64 matrix A, perform C=Not(A), expect UnsupportedType", func(t *testing.T) {
		const (
			rowSize = 5
			colSize = 5
		)
		var (
			A *Matrix[float64]
			C *Matrix[float64]
		)
		t.Run("Setup the matrix", func(t *testing.T) {
			var err error
			if A, err = NewMatrix[float64](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if A == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixFloat64(A, 1)
		})
		t.Run("C=Not(A).  Expect Error", func(t *testing.T) {
			var err error
			if C, err = A.Not(); err == nil {
				t.Fatal("expected UnsupportedType error.  got none.")
			} else {
				if err.Error() != errors.UnsupportedType {
					t.Fatalf("expected UnsupportedType.  Got %v", err)
				}
				if C != nil {
					t.Fatalf("expected nil on error")
				}
			}
		})
	})
}
