package geos

import "testing"

func TestGeosInit(t *testing.T) {
	t.Logf("version: %s", Version())
	Release()
}

func TestGeosWKTCodec(t *testing.T) {
	s := "POINT(123.456 31.654)"
	g := GeometryFromWKT(s)
	tp := g.Type()
	t.Logf("type: %d\nwkt: %s", tp, g.ToWKT())

	Release()
}
