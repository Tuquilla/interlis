package v2_4

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestMultiGeometries(t *testing.T) {
	xtf, _ := os.Open("test.xtf")
	defer xtf.Close()
	decoder := xml.NewDecoder(xtf)
	geometries := ReadGeometry(decoder)

	if len(geometries.MultiSurfaces) != 1 {
		t.Errorf("Not one multipolygon")
	}

	if len(geometries.MultiSurfaces[0].Surfaces) != 2 {
		t.Errorf("Not two polygons inside multipolygon")
	}
}
