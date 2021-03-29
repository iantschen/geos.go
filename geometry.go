package geos

/*
#include <geos_c.h>
*/
import "C"
import "runtime"

type GeometryType int

const (
	Point              GeometryType = C.GEOS_POINT
	LineString         GeometryType = C.GEOS_LINESTRING
	LinearRing         GeometryType = C.GEOS_LINEARRING
	Polygon            GeometryType = C.GEOS_POLYGON
	MultiPoint         GeometryType = C.GEOS_MULTIPOINT
	MultiLineString    GeometryType = C.GEOS_MULTILINESTRING
	MultiPolygon       GeometryType = C.GEOS_MULTIPOLYGON
	GeometryCollection GeometryType = C.GEOS_GEOMETRYCOLLECTION
)

type Geometry struct {
	cval *C.GEOSGeometry
}

func newGeometry(cval *C.GEOSGeometry) *Geometry {
	if cval == nil {
		return nil
	}

	g := &Geometry{
		cval: cval,
	}

	runtime.SetFinalizer(g, func(g *Geometry) {
		C.GEOSGeom_destroy_r(handle, cval)
	})

	return g
}

func (g *Geometry) Envelope() *Geometry {
	v := C.GEOSEnvelope_r(handle, g.cval)
	return newGeometry(v)
}

func (g *Geometry) Buffer(width float64, quadsegs int) *Geometry {
	v := C.GEOSBuffer_r(handle, g.cval, C.double(width), C.int(quadsegs))
	return newGeometry(v)
}

func (g *Geometry) Clone() *Geometry {
	v := C.GEOSGeom_clone_r(handle, g.cval)
	return newGeometry(v)
}

func (g *Geometry) Disjoint(other *Geometry) bool   { return Disjoint(g, other) }
func (g *Geometry) Touches(other *Geometry) bool    { return Touches(g, other) }
func (g *Geometry) Intersects(other *Geometry) bool { return Intersects(g, other) }
func (g *Geometry) Crosses(other *Geometry) bool    { return Crosses(g, other) }
func (g *Geometry) Within(other *Geometry) bool     { return Within(g, other) }
func (g *Geometry) Contains(other *Geometry) bool   { return Contains(g, other) }
func (g *Geometry) Overlaps(other *Geometry) bool   { return Overlaps(g, other) }
func (g *Geometry) Equals(other *Geometry) bool     { return Equals(g, other) }
func (g *Geometry) Covers(other *Geometry) bool     { return Covers(g, other) }
func (g *Geometry) CoveredBy(other *Geometry) bool  { return CoveredBy(g, other) }

func (g *Geometry) IsEmpty() bool {
	return C.GEOSisEmpty_r(handle, g.cval) == 1
}

func (g *Geometry) IsValid() bool {
	return C.GEOSisValid_r(handle, g.cval) == 1
}

func (g *Geometry) Destroy() {
	C.GEOSGeom_destroy_r(handle, g.cval)
}
