package geos

import "testing"

func TestGeosInit(t *testing.T) {
	t.Logf("version: %s", Version())
	Finish()
}
