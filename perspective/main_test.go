package perspective

import (
	"testing"
)

func TestTransform(t *testing.T) {
	p := New(srcPoints, dstPoints)
	x, y := 250.0, 120.0
	rx, ry := p.Transform(x, y)
	ex, ey := 117.27521125839255, 530.9202410878403
	if !eq0(rx, ex) || !eq0(ry, ey) {
		t.Error()
	}
}

func TestTransformInv(t *testing.T) {
	p := New(srcPoints, dstPoints)
	x, y := 130.0, 570.0
	rx, ry := p.TransformInv(x, y)
	ex, ey := 338.99465637447327, 278.6450957956236
	if !eq0(rx, ex) || !eq0(ry, ey) {
		t.Error()
	}
}
