/*
Package perspective is used for creating and applying perspective transforms
with the help of two quadrilaterals. A perspective transform can be used to map
each point of one quadrilateral to another, given the corner coordinates for
the source and destination quadrilaterals.
*/
package perspective

// New takes the corner coordinates of the source and destination quadrilaterals
// in a [x1, y1, x2, y2, x3, y3, x4, y4] form.
func New(srcPoints, dstPoints [8]float64) *Perspective {
	ss := srcPoints[:]
	ds := dstPoints[:]
	c := [9]float64{}
	ci := [9]float64{}
	copy(c[:], getCoeffs(ss, ds, false))
	copy(ci[:], getCoeffs(ss, ds, true))
	return &Perspective{
		SrcPoints: srcPoints,
		DstPoints: dstPoints,
		Coeffs:    c,
		CoeffsInv: ci,
	}
}

type Perspective struct {
	// SrcPoints is an array of the corner coordinates of the source
	// quadrilateral.
	SrcPoints [8]float64
	// DstPoints is an array of the corner coordinates of the destination
	// quadrilateral.
	DstPoints [8]float64
	// Coeffs is a homographic transform matrix, expressed as an array of
	// coefficients.
	Coeffs [9]float64
	// CoeffsInv is an inverse homographic transform matrix, expressed as an
	// array of coefficients.
	CoeffsInv [9]float64
}

// Transform maps a point from the source quadrilateral to the destination
// quadrilateral.
func (t *Perspective) Transform(sx, sy float64) (dx, dy float64) {
	return transform(&t.Coeffs, sx, sy)
}

// TransformInv maps a point from the destination quadrilateral to the source
// quadrilateral.
func (t *Perspective) TransformInv(sx, sy float64) (dx, dy float64) {
	return transform(&t.CoeffsInv, sx, sy)
}

func transform(c *[9]float64, sx, sy float64) (dx, dy float64) {
	dx = (c[0]*sx + c[1]*sy + c[2]) / (c[6]*sx + c[7]*sy + 1)
	dy = (c[3]*sx + c[4]*sy + c[5]) / (c[6]*sx + c[7]*sy + 1)
	return
}
