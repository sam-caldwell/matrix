package matrix

import (
	"testing"
)

func TestMatrix_Rows(t *testing.T) {

	const (
		rowSize = 5
		colSize = 1
	)

	t.Run("test the unexported method rows()", func(t *testing.T) {
		m, err := NewMatrix[int](rowSize, colSize)
		if err != nil {
			t.Fatal(err)
		}
		if len(m.data) != rowSize {
			t.Fatal("dimension failed")
		}
		if rows := m.rows(); rows != rowSize {
			t.Fatal("dimension not working")
		}
	})

	t.Run("test the exported method Rows()", func(t *testing.T) {
		m, err := NewMatrix[int](rowSize, colSize)
		if err != nil {
			t.Fatal(err)
		}
		if len(m.data) != rowSize {
			t.Fatal("dimension failed")
		}
		if rows := m.Rows(); rows != rowSize {
			t.Fatal("dimension not working")
		}
	})

}
