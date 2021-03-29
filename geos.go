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
	handle = C.GEOS_init_r()

	// if handle != nil {
	//     C.GEOSContext_setNoticeHandler_r(handle, C.notice_handler);
	//     C.GEOSContext_setErrorHandler_r(handle, C.error_handler);
	// }
}

// Version returns the version of the GEOS C API in use.
func Version() string {
	return C.GoString(C.GEOSversion())
}

// Error gets the last error that occured in the GEOS C API as a Go error type.
// func Error() error {
// 	return fmt.Errorf("geos: %s", C.GoString(C.gogeos_get_last_error()))
// }

func Finish() {
	C.GEOS_finish_r(handle)
}
