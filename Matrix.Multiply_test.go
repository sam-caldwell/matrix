package matrix

import (
	"github.com/sam-caldwell/errors"
	"testing"
)

func TestMatrix_Multiply(t *testing.T) {

	fillMatrixInt := func(m *Matrix[int], v int, incrementing bool) {
		for r := uint(0); r < m.rows(); r++ {
			for c := uint(0); c < m.cols(); c++ {
				(*m).data[r][c] = v
				if incrementing {
					v++
				}
			}
		}
	}

	fillMatrixFloat := func(m *Matrix[float64], v float64, incrementing bool) {
		for r := uint(0); r < m.rows(); r++ {
			for c := uint(0); c < m.cols(); c++ {
				(*m).data[r][c] = v
				if incrementing {
					v++
				}
			}
		}
	}

	t.Run("Given an integer matrix {A, B}, C=A*B without error", func(t *testing.T) {
		const (
			rowSize  = 5
			colSize  = 5
			valueA   = 2
			valueB   = 2
			expected = 20
		)
		var (
			A *Matrix[int]
			B *Matrix[int]
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
			fillMatrixInt(A, valueA, false)
			if B, err = NewMatrix[int](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixInt(B, valueB, false)
		})
		t.Run("C = A * B", func(t *testing.T) {
			var err error
			if C, err = A.Multiply(B); err != nil {
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

	t.Run("Given an integer matrix {A, B}, C=A*B without error", func(t *testing.T) {
		const (
			valueA   = 2
			valueB   = 0
			expected = 0
		)
		var (
			A *Matrix[int]
			B *Matrix[int]
			C *Matrix[int]
		)
		t.Run("Setup the matrix", func(t *testing.T) {
			var err error
			if A, err = NewMatrix[int](2, 5); err != nil {
				t.Fatal(err)
			}
			if A == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixInt(A, valueA, false)
			if B, err = NewMatrix[int](5, 1); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixInt(B, valueB, false)
		})
		t.Run("C = A * B", func(t *testing.T) {
			var err error
			if C, err = A.Multiply(B); err != nil {
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

	t.Run("Given integer matrix {A,B} of diff sizes, C=A*B expect MatrixDimensionMismatch error", func(t *testing.T) {
		const (
			valueA = 2
			valueB = 2
		)
		var (
			A *Matrix[int]
			B *Matrix[int]
			C *Matrix[int]
		)
		t.Run("Setup the matrix", func(t *testing.T) {
			var err error
			if A, err = NewMatrix[int](5, 10); err != nil {
				t.Fatal(err)
			}
			if A == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixInt(A, valueA, false)
			if B, err = NewMatrix[int](4, 6); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixInt(B, valueB, false)
		})
		t.Run("C = A * B", func(t *testing.T) {
			var err error
			if C, err = A.Multiply(B); err == nil {
				t.Fatal("expected MatrixDimensionMismatch error")
			} else {
				if err.Error() != errors.MatrixDimensionMismatch {
					t.Fatalf("expected MatrixDimensionMismatch.  Got %v", err)
				}
				if C != nil {
					t.Fatalf("on error, expect nil matrix pointer")
				}
			}
		})
	})

	t.Run("Given a float matrix {A, B}, C = A * B without error", func(t *testing.T) {
		const (
			rowSize  = 5
			colSize  = 5
			valueA   = 2.0
			valueB   = 2.0
			expected = 20.0
		)
		var (
			A *Matrix[float64]
			B *Matrix[float64]
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
			fillMatrixFloat(A, valueA, false)
			if B, err = NewMatrix[float64](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixFloat(B, valueB, false)
		})
		t.Run("C = A * B", func(t *testing.T) {
			var err error
			if C, err = A.Multiply(B); err != nil {
				t.Fatal(err)
			}
			for r := uint(0); r < C.rows(); r++ {
				for c := uint(0); c < C.cols(); c++ {
					if actual := (*C).data[r][c]; actual != expected {
						t.Fatalf("value mismatch\n"+
							"actual:   %f\n"+
							"expected: %f", actual, expected)
					}
				}
			}
		})
	})
}
