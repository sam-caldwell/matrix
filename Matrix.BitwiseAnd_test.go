package matrix

import (
	"github.com/sam-caldwell/errors"
	"testing"
)

func TestMatrix_And(t *testing.T) {

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
	t.Run("Given 1's integer matrix A, B, perform C=A AND B without error", func(t *testing.T) {
		const (
			rowSize = 5
			colSize = 5
			valueA  = 1
			valueB  = 1
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
			fillMatrixInt(A, valueA)
			if B, err = NewMatrix[int](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixInt(B, valueB)
		})
		t.Run("C=A AND B.", func(t *testing.T) {
			var err error
			if C, err = A.And(B); err != nil {
				t.Fatal(err)
			}
			for r := uint(0); r < C.rows(); r++ {
				for c := uint(0); c < C.cols(); c++ {
					expected := 1
					if actual := (*C).data[r][c]; actual != expected {
						t.Fatalf("value mismatch\n"+
							"actual:   %d\n"+
							"expected: %d", actual, expected)
					}
				}
			}
		})
	})

	t.Run("Given mixed integer matrix A, B, perform C=A AND B without error", func(t *testing.T) {
		const (
			rowSize = 5
			colSize = 5
			valueA  = 0
			valueB  = 1
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
			fillMatrixInt(A, valueA)
			if B, err = NewMatrix[int](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixInt(B, valueB)
		})
		t.Run("C=A AND B.", func(t *testing.T) {
			var err error
			if C, err = A.And(B); err != nil {
				t.Fatal(err)
			}
			for r := uint(0); r < C.rows(); r++ {
				for c := uint(0); c < C.cols(); c++ {
					expected := 0
					if actual := (*C).data[r][c]; actual != expected {
						t.Fatalf("value mismatch\n"+
							"actual:   %d\n"+
							"expected: %d", actual, expected)
					}
				}
			}
		})
	})

	t.Run("Given float32 matrix A, B, perform C=A AND B and expect UnsupportedType", func(t *testing.T) {
		const (
			rowSize = 5
			colSize = 5
			valueA  = float32(1.0)
			valueB  = float32(2.0)
		)
		var (
			A *Matrix[float32]
			B *Matrix[float32]
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
			if B, err = NewMatrix[float32](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixFloat32(B, valueB)
		})
		t.Run("C=A AND B.  Expect Error", func(t *testing.T) {
			var err error
			if C, err = A.And(B); err == nil {
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

	t.Run("Given float64 matrix A, B, perform C=A AND B and expect UnsupportedType", func(t *testing.T) {
		const (
			rowSize = 5
			colSize = 5
			valueA  = 1.0
			valueB  = 2.0
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
			fillMatrixFloat64(A, 1)
			if B, err = NewMatrix[float64](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixFloat64(B, 2)
		})
		t.Run("C=A AND B.  Expect Error", func(t *testing.T) {
			var err error
			if C, err = A.And(B); err == nil {
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

	t.Run("Given an integer matrix A, B of diff sizes, C=A AND B them, expect error", func(t *testing.T) {
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
			fillMatrixInt(A, 1)
			if B, err = NewMatrix[int](3, 2); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixInt(B, 2)
		})
		t.Run("C=A AND B.", func(t *testing.T) {
			var err error
			if C, err = A.Add(B); err == nil {
				t.Fatalf("expected error but got none")
			} else {
				if err.Error() != errors.MatrixDimensionMismatch {
					t.Fatalf("error mismatch. err: %v", err)
				}
				if C != nil {
					t.Fatalf("expect C is nil on error.")
				}
			}

		})
	})
}
