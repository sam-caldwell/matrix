package matrix

import "testing"

func TestMatrix_type(t *testing.T) {
	t.Run("Matrix should be 2d matrix", func(t *testing.T) {
		m := Matrix[int]{
			data: [][]int{
				{0, 1, 2, 3, 4},
				{0, 1, 2, 3, 4},
			},
		}
		if len(m.data) != 2 {
			t.Fatalf("expected 2 rows")
		}
		if len(m.data[0]) != 5 {
			t.Fatalf("expected 5 cols")
		}
	})
}
