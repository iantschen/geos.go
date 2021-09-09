package geos

/*
#include <geos_c.h>
*/
import "C"

func Intersection(g1, g2 *Geometry) *Geometry {
	v := C.GEOSIntersection_r(handle, g1.cval, g2.cval)
	return newGeometry(v)
}

func Buffer(g *Geometry, width float64, quadsegs int) *Geometry {
	v := C.GEOSBuffer_r(handle, g.cval, C.double(width), C.int(quadsegs))
	return newGeometry(v)
}

func Simplify(g *Geometry, tolerance float64) *Geometry {
	v := C.GEOSSimplify_r(handle, g.cval, C.double(tolerance))
	return newGeometry(v)
}

func TopologyPreserveSimplify(g *Geometry, tolerance float64) *Geometry {
	v := C.GEOSTopologyPreserveSimplify_r(handle, g.cval, C.double(tolerance))
	return newGeometry(v)
}

func PointOnSurface(g *Geometry) *Geometry {
	v := C.GEOSPointOnSurface_r(handle, g.cval)
	return newGeometry(v)
}

func MakeValid(g *Geometry) *Geometry {
	v := C.GEOSMakeValid_r(handle, g.cval)
	return newGeometry(v)
}
