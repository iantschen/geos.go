#include "geos.h"

#define ERROR_LEN 1024

void notice_handler(const char * fmt, ...) {
    va_list ap;
    va_start(ap, fmt);
    fprintf(stderr, "NOTICE: ");
    vfprintf(stderr, fmt, ap);
    va_end(ap);
}

char error_buffer[ERROR_LEN];

void error_handler(const char * fmt, ...) {
    va_list ap;
    va_start(ap, fmt);
    vsnprintf(error_buffer, (size_t) ERROR_LEN, fmt, ap);
    va_end(ap);
}

const char * get_last_error() {
    return error_buffer;
}


GEOSContextHandle_t init() {
    GEOSContextHandle_t h = GEOS_init_r();

	GEOSContext_setNoticeHandler_r(h, notice_handler);
	GEOSContext_setErrorHandler_r(h, error_handler);

    return h;
}