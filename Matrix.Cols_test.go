package matrix

import (
	"testing"
)

func TestMatrix_Cols(t *testing.T) {

	const (
		rowSize = 1
		colSize = 5
	)

	t.Run("test the unexported method cols()", func(t *testing.T) {
		m, err := NewMatrix[int](rowSize, colSize)
		if err != nil {
			t.Fatal(err)
		}
		if len(m.data[0]) != colSize {
			t.Fatal("dimension failed")
		}
		if cols := m.cols(); cols != colSize {
			t.Fatalf("dimension not working: %d", cols)
		}
	})

	t.Run("test the exported method Cols()", func(t *testing.T) {
		m, err := NewMatrix[int](rowSize, colSize)
		if err != nil {
			t.Fatal(err)
		}
		if len(m.data[0]) != colSize {
			t.Fatal("dimension failed")
		}
		if cols := m.Cols(); cols != colSize {
			t.Fatalf("dimension not working: %d", cols)
		}
	})

}
