package geos

/*
#include <geos_c.h>
*/
import "C"

func Disjoint(g1, g2 *Geometry) bool {
	return C.GEOSDisjoint_r(handle, g1.cval, g2.cval) == 1
}

func Touches(g1, g2 *Geometry) bool {
	return C.GEOSTouches_r(handle, g1.cval, g2.cval) == 1
}

func Intersects(g1, g2 *Geometry) bool {
	return C.GEOSIntersects_r(handle, g1.cval, g2.cval) == 1
}

func Crosses(g1, g2 *Geometry) bool {
	return C.GEOSCrosses_r(handle, g1.cval, g2.cval) == 1
}

func Within(g1, g2 *Geometry) bool {
	return C.GEOSWithin_r(handle, g1.cval, g2.cval) == 1
}

func Contains(g1, g2 *Geometry) bool {
	return C.GEOSContains_r(handle, g1.cval, g2.cval) == 1
}

func Overlaps(g1, g2 *Geometry) bool {
	return C.GEOSOverlaps_r(handle, g1.cval, g2.cval) == 1
}

func Equals(g1, g2 *Geometry) bool {
	return C.GEOSEquals_r(handle, g1.cval, g2.cval) == 1
}

func Covers(g1, g2 *Geometry) bool {
	return C.GEOSCovers_r(handle, g1.cval, g2.cval) == 1
}

func CoveredBy(g1, g2 *Geometry) bool {
	return C.GEOSCoveredBy_r(handle, g1.cval, g2.cval) == 1
}

func IsEmpty(g *Geometry) bool {
	return C.GEOSisEmpty_r(handle, g.cval) == 1
}

func IsValid(g *Geometry) bool {
	return C.GEOSisValid_r(handle, g.cval) == 1
}

func boolFrom(v C.char) bool {
	return v == 1
}
