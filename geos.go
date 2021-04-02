package geos

/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/

import "C"

var (
	handle C.GEOSContextHandle_t
	reader *C.GEOSWKTReader
	writer *C.GEOSWKTWriter
)

func init() {
	handle = C.newHandle()
	reader = C.newWKTReader(handle)
	writer = C.newWKTWriter(handle)
}

// Version returns the version of the GEOS C API in use.
func Version() string {
	return C.GoString(C.GEOSversion())
}

func Release() {
	C.GEOSWKTReader_destroy_r(handle, reader)
	C.GEOSWKTWriter_destroy_r(handle, writer)
	C.GEOS_finish_r(handle)
}
