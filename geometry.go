package geos

/*
#include <geos_c.h>
*/
import "C"

type Geometry struct {
	h *C.GEOSGeometry
}

// func NewGeometry() *Geometry {

// }

// func (g *Geometry) Clone() *Geometry {

// }

func (g *Geometry) GEOSDisjoint(other *Geometry) bool   { return GEOSDisjoint(g, other) }
func (g *Geometry) GEOSTouches(other *Geometry) bool    { return GEOSTouches(g, other) }
func (g *Geometry) GEOSIntersects(other *Geometry) bool { return GEOSINTERsects(g, other) }
func (g *Geometry) GEOSCrosses(other *Geometry) bool    { return GEOSCrosses(g, other) }
func (g *Geometry) GEOSWithin(other *Geometry) bool     { return GEOSWithin(g, other) }
func (g *Geometry) GEOSContains(other *Geometry) bool   { return GEOSContains(g, other) }
func (g *Geometry) GEOSOverlaps(other *Geometry) bool   { return GEOSOverlaps(g, other) }
func (g *Geometry) GEOSEquals(other *Geometry) bool     { return GEOSEquals(g, other) }
func (g *Geometry) GEOSCovers(other *Geometry) bool     { return GEOSCovers(g, other) }
func (g *Geometry) GEOSCoveredBy(other *Geometry) bool  { return GEOSCoveredBy(g, other) }

func (g *Geometry) Destroy() {
	C.GEOSGeom_destroy_r(handle, g.h)
}

func GEOSDisjoint(g1, g2 *Geometry) bool  { return false }
func GEOSTouches(g1, g2 *Geometry) bool   { return false }
func GEOSIntersect(g1, g2 *Geometry) bool { return false }
func GEOSCrosses(g1, g2 *Geometry) bool   { return false }
func GEOSWithin(g1, g2 *Geometry) bool    { return false }
func GEOSContains(g1, g2 *Geometry) bool  { return false }
func GEOSOverlaps(g1, g2 *Geometry) bool  { return false }
func GEOSEquals(g1, g2 *Geometry) bool    { return false }
func GEOSCovers(g1, g2 *Geometry) bool    { return false }
func GEOSCoveredBy(g1, g2 *Geometry) bool { return false }
