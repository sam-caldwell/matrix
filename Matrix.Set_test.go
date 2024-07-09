package matrix

import (
	"github.com/sam-caldwell/errors"
	"testing"
)

func TestMatrix_Set(t *testing.T) {
	t.Run("Write to empty matrix.  Expect BoundsCheckError", func(t *testing.T) {
		m := &(Matrix[int]{})
		if len(m.data) != 0 {
			t.Fatal("expect empty matrix")
		}
		if err := m.Set(0, 0, 5); err == nil {
			t.Fatal("expected an error. got none")
		} else {
			if err.Error() != errors.BoundsCheckError {
				t.Fatalf("expected BoundsCheckError.  Got %v", err)
			}
		}
	})
	t.Run("Write to 5x5 matrix (int).  Expect no error.  Read and verify", func(t *testing.T) {
		var (
			err error
			m   *Matrix[int]
			v   int
		)
		if m, err = NewMatrix[int](5, 5); err != nil {
			t.Fatal(err)
		}
		if err = m.Set(3, 2, -99); err != nil {
			t.Fatal(err)
		}
		if v, err = m.Get(3, 2); err != nil {
			t.Fatal(err)
		} else {
			if v != -99 {
				t.Fatalf("value mismatch.  wanted -99.  got: %d", v)
			}
		}
	})
	t.Run("Write out of bounds given a 5x5 matrix, expect BoundsCheckError", func(t *testing.T) {
		var (
			err error
			m   *Matrix[int]
		)
		if m, err = NewMatrix[int](5, 5); err != nil {
			t.Fatal(err)
		}
		if err = m.Set(10, 54, 1337); err == nil {
			t.Fatal("expected error not found")
		} else {
			if err.Error() != errors.BoundsCheckError {
				t.Fatalf("expected BoundsCheckError.  Got: %v", err)
			}
		}
	})

	t.Run("Write to 5x5 matrix (float64).  Expect no error.  Read and verify", func(t *testing.T) {
		const pi = 3.1415
		var (
			err error
			m   *Matrix[float64]
			v   float64
		)
		if m, err = NewMatrix[float64](5, 5); err != nil {
			t.Fatal(err)
		}
		if err = m.Set(3, 2, pi); err != nil {
			t.Fatal(err)
		}
		if v, err = m.Get(3, 2); err != nil {
			t.Fatal(err)
		} else {
			if v != pi {
				t.Fatalf("value mismatch.  wanted pi.  got: %f", v)
			}
		}
	})
}
