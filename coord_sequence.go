package geos

/*
#include "geos.h"
*/
import "C"

type CoordSequence struct {
	ptr *C.GEOSCoordSequence
}
