package geos

/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"

var (
	handle C.GEOSContextHandle_t
)

func init() {
	handle = C.newHandle()
}

// Version returns the version of the GEOS C API in use.
func Version() string {
	return C.GoString(C.GEOSversion())
}

func Release() {
	C.GEOS_finish_r(handle)
}
