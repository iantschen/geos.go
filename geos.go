package geos

/*
#cgo LDFLAGS: -lgeos_c
#include "geos.h"
*/
import "C"
import "fmt"

var (
	handle C.GEOSContextHandle_t
)

func init() {
	handle = C.init()
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
	C.GEOS_finish_r(handle)
}
