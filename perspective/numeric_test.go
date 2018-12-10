package perspective

import (
	"testing"
)

func TestTranspose(t *testing.T) {
	res := transpose(a)
	exp := [][]float64{
		[]float64{0, 1, 2, -1, 0, 1, 2, -1},
		[]float64{1, 2, 3, -2, 1, 2, 3, -2},
		[]float64{2, 3, 4, -3, 2, 3, 4, -3},
		[]float64{3, 4, 5, -4, 3, 4, 5, -4},
		[]float64{4, 5, 6, -5, 4, 5, 6, -5},
		[]float64{5, 6, 7, -6, 5, 6, 7, -6},
		[]float64{6, 7, 8, -7, 6, 7, 8, -7},
		[]float64{7, 8, 9, -8, 7, 8, 9, -8},
	}
	if !eq2(res, exp) {
		t.Error()
	}
}

func TestDotMMSlmall(t *testing.T) {
	res := dotMMSmall(a, b)
	exp := [][]float64{
		[]float64{-41, 2, -56, 54, -31, -6, -83, 50},
		[]float64{-50, 3, -72, 68, -45, -3, -108, 68},
		[]float64{-59, 4, -88, 82, -59, 0, -133, 86},
		[]float64{50, -3, 72, -68, 45, 3, 108, -68},
		[]float64{-41, 2, -56, 54, -31, -6, -83, 50},
		[]float64{-50, 3, -72, 68, -45, -3, -108, 68},
		[]float64{-59, 4, -88, 82, -59, 0, -133, 86},
		[]float64{50, -3, 72, -68, 45, 3, 108, -68},
	}
	if !eq2(res, exp) {
		t.Error()
	}
}

func TestInv_invertible(t *testing.T) {
	res := inv(invertible)
	exp := [][]float64{
		[]float64{-1.5, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, -1},
		[]float64{0.5, -1, 0, 0, 0, 0, 0, 0.5},
		[]float64{0.5, 0, -1, 0, 0, 0, 0, 0.5},
		[]float64{0.5, 0, 0, -1, 0, 0, 0, 0.5},
		[]float64{0.5, 0, 0, 0, -1, 0, 0, 0.5},
		[]float64{0.5, 0, 0, 0, 0, -1, 0, 0.5},
		[]float64{0.5, 0, 0, 0, 0, 0, -1, 0.5},
		[]float64{-1, 0.5, 0.5, 0.5, 0.5, 0.5, 0.5, -1.5},
	}
	if !eq2(res, exp) {
		t.Error()
	}
}

func TestInv_singular(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error()
		}
	}()
	inv(a)
}

func TestDotMV(t *testing.T) {
	res := dotMV(a, points)
	exp := []float64{168, 204, 240, -204, 168, 204, 240, -204}
	if !eq1(res, exp) {
		t.Error()
	}
}

func TestGetCoeffs(t *testing.T) {
	res := getCoeffs(srcPoints[:], dstPoints[:], false)
	exp := []float64{
		0.3869749384, 0.0426817448, 59.2427947969,
		0.9589610618, 0.4562821238, 434.8644299345,
		0.0012901794, 0.0004268174, 1,
	}
	if !eq1(res, exp) {
		t.Error()
	}
}

func TestGetCoeffs_inv(t *testing.T) {
	res := getCoeffs(srcPoints[:], dstPoints[:], true)
	exp := []float64{
		1.9955408809, -0.1282507787, -62.4497171511,
		-2.9335671323, 2.2894572644, -821.8108124927,
		-0.0013225082, -0.0008117138, 1,
	}
	if !eq1(res, exp) {
		t.Error()
	}
}
