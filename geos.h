#include <stdlib.h>
#include <stdio.h>
#include <stdarg.h>

//#include <string.h>

#include <geos_c.h>

// void notice_handler(const char *fmt, ...);
// void error_handler(const char *fmt, ...);

GEOSContextHandle_t newHandle();
GEOSWKTReader * newWKTReader(GEOSContextHandle_t h);
GEOSWKTWriter * newWKTWriter(GEOSContextHandle_t h);
