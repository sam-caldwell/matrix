package matrix

import (
	"github.com/sam-caldwell/errors"
	"testing"
)

func TestMatrix_Add(t *testing.T) {

	fillMatrixInt := func(m *Matrix[int], v int) {
		i := v
		for r := uint(0); r < m.rows(); r++ {
			for c := uint(0); c < m.cols(); c++ {
				(*m).data[r][c] = v
				i++
			}
		}
	}
	fillMatrixFloat := func(m *Matrix[float64], v float64) {
		i := v
		for r := uint(0); r < m.rows(); r++ {
			for c := uint(0); c < m.cols(); c++ {
				(*m).data[r][c] = v
				i++
			}
		}
	}
	t.Run("Given an integer matrix A and B, add A+B to produce C without error", func(t *testing.T) {
		const (
			rowSize = 5
			colSize = 5
			valueA  = 1
			valueB  = 2
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
		t.Run("Add A+B to create Matrix C", func(t *testing.T) {
			var err error
			if C, err = A.Add(B); err != nil {
				t.Fatal(err)
			}
			a := valueA
			b := valueB
			for r := uint(0); r < C.rows(); r++ {
				for c := uint(0); c < C.cols(); c++ {
					expected := a + b
					if actual := (*C).data[r][c]; actual != expected {
						t.Fatalf("value mismatch\n"+
							"actual:   %d\n"+
							"expected: %d", actual, expected)
					}

				}
			}
		})
	})

	t.Run("Given a float matrix A and B, add A+B to produce C without error", func(t *testing.T) {
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
			fillMatrixFloat(A, 1)
			if B, err = NewMatrix[float64](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixFloat(B, 2)
		})
		t.Run("Add A+B to create Matrix C", func(t *testing.T) {
			var err error
			if C, err = A.Add(B); err != nil {
				t.Fatal(err)
			}
			a := valueA
			b := valueB
			for r := uint(0); r < C.rows(); r++ {
				for c := uint(0); c < C.cols(); c++ {
					expected := a + b
					if actual := (*C).data[r][c]; actual != expected {
						t.Fatalf("value mismatch\n"+
							"actual:   %f\n"+
							"expected: %f", actual, expected)
					}

				}
			}
		})
	})

	t.Run("Given an integer matrix A and B of diff sizes, add them and expect error", func(t *testing.T) {
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
		t.Run("Add A+B to create Matrix C", func(t *testing.T) {
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
