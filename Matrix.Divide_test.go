package matrix

import (
	"github.com/sam-caldwell/errors"
	"testing"
)

func TestMatrix_Divide(t *testing.T) {

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

	t.Run("Given an integer matrix {A, B}, C=A/B without error", func(t *testing.T) {
		const (
			rowSize  = 5
			colSize  = 5
			valueA   = 2
			valueB   = 2
			expected = 1
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
			fillMatrixInt(A, valueA, true)
			if B, err = NewMatrix[int](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixInt(B, valueB, true)
		})
		t.Run("C = A / B", func(t *testing.T) {
			var err error
			if C, err = A.Divide(B); err != nil {
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

	t.Run("Given an integer matrix {A, B}, C=A/B to cause DivideByZero error", func(t *testing.T) {
		const (
			rowSize = 5
			colSize = 5
			valueA  = 2
			valueB  = 0
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
			fillMatrixInt(A, valueA, true)
			if B, err = NewMatrix[int](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixInt(B, valueB, false)
		})
		t.Run("C = A / B", func(t *testing.T) {
			var err error
			if C, err = A.Divide(B); err == nil {
				t.Fatal("expected DivisionByZero error")
			} else {
				if err.Error() != errors.DivisionByZero {
					t.Fatalf("expected DivisionByZero.  Got %v", err)
				}
				if C != nil {
					t.Fatalf("on division by zero error, expect nil matrix pointer")
				}
			}
		})
	})

	t.Run("Given a float matrix {A, B}, C=A/B without error", func(t *testing.T) {
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
			fillMatrixFloat(A, valueA, true)
			if B, err = NewMatrix[float64](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixFloat(B, valueB, true)
		})
		t.Run("C = A / B", func(t *testing.T) {
			var err error
			if C, err = A.Divide(B); err != nil {
				t.Fatal(err)
			}
			a := valueA
			b := valueB
			for r := uint(0); r < C.rows(); r++ {
				for c := uint(0); c < C.cols(); c++ {
					expected := a / b
					if actual := (*C).data[r][c]; actual != expected {
						t.Fatalf("value mismatch\n"+
							"actual:   %f\n"+
							"expected: %f", actual, expected)
					}
					a++
					b++
				}
			}
		})
	})

	t.Run("Given a float matrix {A, B}, C=A/B to cause DivideByZero error", func(t *testing.T) {
		const (
			rowSize  = 5
			colSize  = 5
			valueA   = 1.0
			valueB   = 0
			expected = float64(0)
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
			fillMatrixFloat(A, valueA, true)
			if B, err = NewMatrix[float64](rowSize, colSize); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixFloat(B, valueB, false)
		})
		t.Run("C = A / B", func(t *testing.T) {
			var err error
			if C, err = A.Divide(B); err == nil {
				t.Fatal("expected DivisionByZero error")
			} else {
				if err.Error() != errors.DivisionByZero {
					t.Fatalf("expected DivisionByZero.  Got %v", err)
				}
				if C != nil {
					t.Fatalf("on division by zero error, expect nil matrix pointer")
				}
			}
		})
	})

	t.Run("Given an integer matrix {A, B} of diff sizes, divide them and expect error", func(t *testing.T) {
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
			fillMatrixInt(A, 1, true)
			if B, err = NewMatrix[int](3, 2); err != nil {
				t.Fatal(err)
			}
			if B == nil {
				t.Fatal("unexpected nil matrix")
			}
			fillMatrixInt(B, 2, true)
		})
		t.Run("C = A / B", func(t *testing.T) {
			var err error
			if C, err = A.Divide(B); err == nil {
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
