package v2_4

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestMultiPolygon(t *testing.T) {
	xtf, _ := os.Open("test_2_4.xtf")
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

func TestMultiLine(t *testing.T) {
	xtf, _ := os.Open("test_2_4.xtf")
	defer xtf.Close()
	decoder := xml.NewDecoder(xtf)
	geometries := ReadGeometry(decoder)

	if len(geometries.MutliPolylines) != 1 {
		t.Errorf("Not one multiline")
	}

	if len(geometries.MutliPolylines[0].Polylines) != 2 {
		t.Errorf("Not two lines inside multiline")
	}
}

func TestMultiPoint(t *testing.T) {
	xtf, _ := os.Open("test_2_4.xtf")
	defer xtf.Close()
	decoder := xml.NewDecoder(xtf)
	geometries := ReadGeometry(decoder)

	if len(geometries.MultiCoords) != 2 {
		t.Errorf("Not tow multipoint")
	}

	if len(geometries.MultiCoords[0].Coords) != 4 {
		t.Errorf("Not four points inside multipoint")
	}
}
