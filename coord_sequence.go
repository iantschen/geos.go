package geos

/*
#include "geos.h"
*/
import "C"
import "runtime"

type CoordSequence struct {
	cval *C.GEOSCoordSequence
}

func NewCoordSequence(coords []Coordinate) *CoordSequence {
	v := C.GEOSCoordSeq_create_r(handle, C.uint(len(coords)), 2)
	if v == nil {
		return nil
	}

	cs := &CoordSequence{cval: v}
	for i, c := range coords {
		C.GEOSCoordSeq_setX_r(handle, v, C.uint(i), C.double(c[0]))
		C.GEOSCoordSeq_setY_r(handle, v, C.uint(i), C.double(c[1]))
	}

	runtime.SetFinalizer(cs, func(*CoordSequence) {
		C.GEOSCoordSeq_destroy_r(handle, v)
	})

	return cs
}

func (cs *CoordSequence) Clone() *CoordSequence {
	v := C.GEOSCoordSeq_clone_r(handle, cs.cval)
	if v == nil {
		return nil
	}

	cs1 := &CoordSequence{cval: v}
	runtime.SetFinalizer(cs, func(*CoordSequence) {
		C.GEOSCoordSeq_destroy_r(handle, v)
	})

	return cs1
}
