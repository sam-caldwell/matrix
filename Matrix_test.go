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
	t.Run("Matrix should support int", func(t *testing.T) {
		_ = Matrix[int]{
			data: [][]int{
				{0, 1},
				{0, 1},
			},
		}

	})
	t.Run("Matrix should support int8", func(t *testing.T) {
		_ = Matrix[int8]{
			data: [][]int8{
				{0, 1},
				{0, 1},
			},
		}

	})
	t.Run("Matrix should support int16", func(t *testing.T) {
		_ = Matrix[int16]{
			data: [][]int16{
				{0, 1},
				{0, 1},
			},
		}

	})
	t.Run("Matrix should support int32", func(t *testing.T) {
		_ = Matrix[int32]{
			data: [][]int32{
				{0, 1},
				{0, 1},
			},
		}

	})
	t.Run("Matrix should support int64", func(t *testing.T) {
		_ = Matrix[int64]{
			data: [][]int64{
				{0, 1},
				{0, 1},
			},
		}

	})
	t.Run("Matrix should support uint", func(t *testing.T) {
		_ = Matrix[uint]{
			data: [][]uint{
				{0, 1},
				{0, 1},
			},
		}

	})
	t.Run("Matrix should support uint8", func(t *testing.T) {
		_ = Matrix[uint8]{
			data: [][]uint8{
				{0, 1},
				{0, 1},
			},
		}

	})
	t.Run("Matrix should support uint16", func(t *testing.T) {
		_ = Matrix[uint16]{
			data: [][]uint16{
				{0, 1},
				{0, 1},
			},
		}

	})
	t.Run("Matrix should support uint32", func(t *testing.T) {
		_ = Matrix[uint32]{
			data: [][]uint32{
				{0, 1},
				{0, 1},
			},
		}

	})
	t.Run("Matrix should support uint64", func(t *testing.T) {
		_ = Matrix[uint64]{
			data: [][]uint64{
				{0, 1},
				{0, 1},
			},
		}

	})
	t.Run("Matrix should support float32", func(t *testing.T) {
		_ = Matrix[float32]{
			data: [][]float32{
				{0.001, 0.002},
				{0.003, 0.004},
			},
		}

	})
	t.Run("Matrix should support float64", func(t *testing.T) {
		_ = Matrix[float64]{
			data: [][]float64{
				{0.001, 0.002},
				{0.003, 0.004},
			},
		}

	})
}
