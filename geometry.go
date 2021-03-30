package geos

/*
#include <geos_c.h>
#include <stdlib.h>
*/
import "C"
import (
	"runtime"
	"unsafe"
)

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

func GeometryFromWKT(wkt string) *Geometry {
	cstr := C.CString(wkt)
	defer C.free(unsafe.Pointer(cstr))

	v := C.GEOSWKTReader_read_r(handle, reader, cstr)
	return newGeometry(v)
}

func NewPoint(c *Coordinate) (*Geometry, error) {

}

func NewLineString() (*Geometry, error) {

}

func NewLinearRing() (*Geometry, error) {

}

func NewPolygon(shell *Geometry, holes []*Geometry) (*Geometry, error) {

}

func NewMultiPoint() (*Geometry, error) {

}

func NewMultiLineString() (*Geometry, error) {

}

func NewMultiPolygon() (*Geometry, error) {

}

func NewGeometryCollection() (*Geometry, error) {

}

func (g *Geometry) MinX() float64 {
	var mx C.double
	C.GEOSGeom_getXMin_r(handle, g.cval, &mx)
	return float64(mx)
}

func (g *Geometry) MaxX() float64 {
	var mx C.double
	C.GEOSGeom_getXMax_r(handle, g.cval, &mx)
	return float64(mx)
}

func (g *Geometry) MinY() float64 {
	var my C.double
	C.GEOSGeom_getYMin_r(handle, g.cval, &my)
	return float64(my)
}

func (g *Geometry) MaxY() float64 {
	var my C.double
	C.GEOSGeom_getYMax_r(handle, g.cval, &my)
	return float64(my)
}

func (g *Geometry) Type() GeometryType {
	id := C.GEOSGeomTypeId_r(handle, g.cval)
	return GeometryType(id)
}

func (g *Geometry) Envelope() *Geometry {
	v := C.GEOSEnvelope_r(handle, g.cval)
	return newGeometry(v)
}

func (g *Geometry) ConvexHull() *Geometry {
	v := C.GEOSConvexHull_r(handle, g.cval)
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

func (g *Geometry) ToWKT() string {
	v := C.GEOSWKTWriter_write_r(handle, writer, g.cval)
	if v != nil {
		defer C.free(unsafe.Pointer(v))
	}
	return C.GoString(v)
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
