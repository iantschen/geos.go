#include <stdarg.h>

#include "geos.h"

void notice_handler(const char * fmt, ...)
{
    va_list ap;
    va_start(ap, fmt);
    fprintf(stderr, "NOTICE: ");
    vfprintf(stderr, fmt, ap);
    fprintf(stderr, "\n");
    va_end(ap);
}

void error_handler(const char * fmt, ...)
{
    va_list ap;
    va_start(ap, fmt);
    fprintf(stderr, "ERROR: ");
    vfprintf(stderr, fmt, ap);
    fprintf(stderr, "\n");
    va_end(ap);
}

GEOSContextHandle_t newHandle() 
{
    GEOSContextHandle_t h = GEOS_init_r();

	GEOSContext_setNoticeHandler_r(h, notice_handler);
	GEOSContext_setErrorHandler_r(h, error_handler);

    return h;
}

GEOSWKTReader * newWKTReader(GEOSContextHandle_t h) 
{
    return GEOSWKTReader_create_r(h);
}

GEOSWKTWriter * newWKTWriter(GEOSContextHandle_t h) 
{
    return GEOSWKTWriter_create_r(h);
}

GEOSWKBReader * newWKBReader(GEOSContextHandle_t h) 
{
    return GEOSWKBReader_create_r(h);
}

GEOSWKBWriter * newWKBWriter(GEOSContextHandle_t h) 
{
    return GEOSWKBWriter_create_r(h);
}