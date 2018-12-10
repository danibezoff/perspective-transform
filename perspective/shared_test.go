package perspective

import "math"

var (
	a = [][]float64{
		[]float64{0, 1, 2, 3, 4, 5, 6, 7},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8},
		[]float64{2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{-1, -2, -3, -4, -5, -6, -7, -8},
		[]float64{0, 1, 2, 3, 4, 5, 6, 7},
		[]float64{1, 2, 3, 4, 5, 6, 7, 8},
		[]float64{2, 3, 4, 5, 6, 7, 8, 9},
		[]float64{-1, -2, -3, -4, -5, -6, -7, -8},
	}
	b = [][]float64{
		[]float64{-1, 1, -2, 2, -3, 3, -4, 4},
		[]float64{-1, 0, -2, 0, -3, 0, -4, 0},
		[]float64{-2, 1, -2, 2, -2, 3, -2, 4},
		[]float64{1, -1, -2, 2, -3, 0, -3, 4},
		[]float64{0, -1, -2, 2, -3, 0, -3, 4},
		[]float64{-1, 0, -2, 2, 0, -3, -3, 0},
		[]float64{-1, 0, -2, 2, 3, -3, -3, 0},
		[]float64{-4, 1, -2, 2, -3, 3, -3, 2},
	}
	invertible = [][]float64{
		[]float64{0, 1, 1, 1, 1, 1, 1, 2},
		[]float64{1, 0, 1, 1, 1, 1, 1, 1},
		[]float64{1, 1, 0, 1, 1, 1, 1, 1},
		[]float64{1, 1, 1, 0, 1, 1, 1, 1},
		[]float64{1, 1, 1, 1, 0, 1, 1, 1},
		[]float64{1, 1, 1, 1, 1, 0, 1, 1},
		[]float64{1, 1, 1, 1, 1, 1, 0, 1},
		[]float64{2, 1, 1, 1, 1, 1, 1, 0},
	}
	points    = []float64{1, 2, 3, 4, 5, 6, 7, 8}
	srcPoints = [8]float64{158, 64, 494, 69, 495, 404, 158, 404}
	dstPoints = [8]float64{100, 500, 152, 564, 148, 604, 100, 560}
)

const e = 1e-5

func eq0(a, b float64) bool {
	if math.Abs(a-b) > e {
		return false
	}
	return true
}

func eq1(a, b []float64) bool {
	for i := range a {
		if math.Abs(a[i]-b[i]) > e {
			return false
		}
	}
	return true
}

func eq2(a, b [][]float64) bool {
	for i := range a {
		for j := range a[0] {
			if math.Abs(a[i][j]-b[i][j]) > e {
				return false
			}
		}
	}
	return true
}
