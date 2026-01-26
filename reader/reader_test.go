package interlis

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"

	interlis2 "github.com/tuquilla/interlis"
)

func TestXtfReader(t *testing.T) {

	// Prepare file
	xtf, _ := os.Open("test.xtf")
	defer xtf.Close()
	decoder := xml.NewDecoder(xtf)

	geometries := interlis2.ReadGeometry(decoder)
	if geometries.Surfaces[0].Boundary.Polyline.Coords[0].C1 != "1" {
		t.Errorf("C1 was not 1")
	}
	if geometries.Surfaces[0].Boundary.Polyline.Coords[0].C2 != "2" {
		t.Errorf("C2 was not 2")
	}
	if geometries.Surfaces[0].Boundary.Polyline.Coords[0].C3 != "" {
		t.Errorf("C3 was not empty")
	}
}

func TestMultiGeometries(t *testing.T) {
	xtf, _ := os.Open("test_multi.xtf")
	defer xtf.Close()
	decoder := xml.NewDecoder(xtf)

	geometries := interlis2.ReadGeometry(decoder)
	fmt.Println(geometries)
	fmt.Println(geometries)
	if len(geometries.MultiSurfaces) != 1 {
		t.Errorf("Not one multipolygon")
	}

	if len(geometries.MultiSurfaces[0].Surfaces) != 2 {
		t.Errorf("Not two polygons inside multipolygon")
	}
}
