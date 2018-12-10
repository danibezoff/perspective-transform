// This is intended to work only with 8x8 matrices

package perspective

import (
	"math"
)

func transpose(x [][]float64) [][]float64 {
	res := make([][]float64, 8)
	for i := range x {
		res[i] = make([]float64, 8)
		for j := range x[i] {
			res[i][j] = x[j][i]
		}
	}
	return res
}

func dotMMSmall(x, y [][]float64) [][]float64 {
	res := newMat()
	for i := 7; i >= 0; i-- {
		bar := x[i]
		for k := 7; k >= 0; k-- {
			woo := bar[7] * y[7][k]
			for j := 6; j >= 1; j -= 2 {
				i0 := j - 1
				woo += bar[j]*y[j][k] + bar[i0]*y[i0][k]
			}
			woo += bar[0] * y[0][k]
			res[i][k] = woo
		}
	}
	return res
}

func dotMV(x [][]float64, y []float64) []float64 {
	res := make([]float64, 8)
	for i := 7; i >= 0; i-- {
		res[i] = dotVV(x[i], y)
	}
	return res
}

func dotVV(x, y []float64) float64 {
	res := x[7] * y[7]
	for i := 6; i >= 1; i -= 2 {
		i1 := i - 1
		res += x[i]*y[i] + x[i1]*y[i1]
	}
	res += x[0] * y[0]
	return res
}

func inv(x [][]float64) [][]float64 {
	A := copyMat(x)
	I := identity()
	for j := 0; j < 8; j++ {
		i0 := -1
		v0 := -1.0
		for i := j; i != 8; i++ {
			if k := math.Abs(A[i][j]); k > v0 {
				i0 = i
				v0 = k
			}
		}
		Aj := A[i0]
		A[i0] = A[j]
		A[j] = Aj
		Ij := I[i0]
		I[i0] = I[j]
		I[j] = Ij
		foo := Aj[j]
		for k := j; k != 8; k++ {
			Aj[k] /= foo
		}
		for k := 7; k != -1; k-- {
			Ij[k] /= foo
		}
		for i := 7; i != -1; i-- {
			if i == j {
				continue
			}
			Ai := A[i]
			Ii := I[i]
			bar := Ai[j]
			for k := j + 1; k != 8; k++ {
				Ai[k] -= Aj[k] * bar
			}
			for k := 7; k > 0; k-- {
				Ii[k] -= Ij[k] * bar
				k--
				Ii[k] -= Ij[k] * bar
			}
		}
	}
	return I
}

func copyMat(x [][]float64) [][]float64 {
	res := newMat()
	for i := range res {
		copy(res[i], x[i])
	}
	return res
}

func identity() [][]float64 {
	return [][]float64{
		[]float64{1, 0, 0, 0, 0, 0, 0, 0},
		[]float64{0, 1, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 1, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 1, 0, 0, 0, 0},
		[]float64{0, 0, 0, 0, 1, 0, 0, 0},
		[]float64{0, 0, 0, 0, 0, 1, 0, 0},
		[]float64{0, 0, 0, 0, 0, 0, 1, 0},
		[]float64{0, 0, 0, 0, 0, 0, 0, 1},
	}
}

func newMat() [][]float64 {
	res := make([][]float64, 8)
	for i := range res {
		res[i] = make([]float64, 8)
	}
	return res
}

func getCoeffs(srcPts, dstPts []float64, inv bool) []float64 {
	if inv {
		srcPts, dstPts = dstPts, srcPts
	}
	r1 := []float64{srcPts[0], srcPts[1], 1, 0, 0, 0, -1 * dstPts[0] * srcPts[0], -1 * dstPts[0] * srcPts[1]}
	r2 := []float64{0, 0, 0, srcPts[0], srcPts[1], 1, -1 * dstPts[1] * srcPts[0], -1 * dstPts[1] * srcPts[1]}
	r3 := []float64{srcPts[2], srcPts[3], 1, 0, 0, 0, -1 * dstPts[2] * srcPts[2], -1 * dstPts[2] * srcPts[3]}
	r4 := []float64{0, 0, 0, srcPts[2], srcPts[3], 1, -1 * dstPts[3] * srcPts[2], -1 * dstPts[3] * srcPts[3]}
	r5 := []float64{srcPts[4], srcPts[5], 1, 0, 0, 0, -1 * dstPts[4] * srcPts[4], -1 * dstPts[4] * srcPts[5]}
	r6 := []float64{0, 0, 0, srcPts[4], srcPts[5], 1, -1 * dstPts[5] * srcPts[4], -1 * dstPts[5] * srcPts[5]}
	r7 := []float64{srcPts[6], srcPts[7], 1, 0, 0, 0, -1 * dstPts[6] * srcPts[6], -1 * dstPts[6] * srcPts[7]}
	r8 := []float64{0, 0, 0, srcPts[6], srcPts[7], 1, -1 * dstPts[7] * srcPts[6], -1 * dstPts[7] * srcPts[7]}
	matA := [][]float64{r1, r2, r3, r4, r5, r6, r7, r8}
	matB := dstPts

	matC, err := getMatC(matA)
	if err == true {
		return []float64{1, 0, 0, 0, 1, 0, 0, 0, 1}
	}

	matD := dotMMSmall(matC, transpose(matA))
	matX := dotMV(matD, matB)
	matX = append(matX, 1)
	return matX
}

func getMatC(matA [][]float64) (res [][]float64, err bool) {
	defer func() {
		if r := recover(); r != nil {
			res, err = nil, true
		}
	}()
	return inv(dotMMSmall(transpose(matA), matA)), false
}
