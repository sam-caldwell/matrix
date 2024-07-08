package matrix

import "testing"

func TestMatrix_type(t *testing.T) {
	t.Run("Matrix should be 2d matrix", func(t *testing.T) {
		m := Matrix[int]{
			{0, 1, 2, 3, 4},
			{0, 1, 2, 3, 4},
		}
		if len(m) != 5 || len(m[0]) != 2 {
			t.Fatalf("expect 2d matrix")
		}
	})

}
