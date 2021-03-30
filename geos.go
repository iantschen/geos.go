package geos

/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"
import "fmt"

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

// Error gets the last error that occured in the GEOS C API as a Go error type.
func Error() error {
	return fmt.Errorf("geos: %s", C.GoString(C.get_last_error()))
}

func Release() {
	C.GEOSWKTReader_destroy_r(handle, reader)
	C.GEOSWKTWriter_destroy_r(handle, writer)
	C.GEOS_finish_r(handle)
}
