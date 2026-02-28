package v2_3

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestXtfReader(t *testing.T) {

	// Prepare file
	xtf, _ := os.Open("test.xtf")
	defer xtf.Close()
	decoder := xml.NewDecoder(xtf)

	geometries := ReadGeometry(decoder)

	// Point
	if geometries.Coords[0].C1 != "2718805.482" {
		t.Errorf("C1 was incorrect, C1 was %s", geometries.Coords[0].C1)
	}
	if geometries.Coords[0].C2 != "1228925.698" {
		t.Errorf("C2 was incorrect, C1 was %s", geometries.Coords[0].C2)
	}
	if geometries.Coords[0].C3 != "" {
		t.Errorf("C3 was incorrect, C1 was %s", geometries.Coords[0].C3)
	}

	// Line
	if geometries.Polylines[0].Coords[0].C1 != "2718802.115" {
		t.Errorf("C1 was incorrect, C1 was %s", geometries.Polylines[0].Coords[0].C1)
	}
	if geometries.Polylines[0].Coords[0].C2 != "1228947.023" {
		t.Errorf("C2 was incorrect, C1 was %s", geometries.Polylines[0].Coords[0].C2)
	}
	if geometries.Polylines[0].Coords[0].C3 != "" {
		t.Errorf("C3 was incorrect, C1 was %s", geometries.Polylines[0].Coords[0].C3)
	}

	// Polygon (surface)
	if geometries.Surfaces[0].Boundary[0].Polyline.Coords[0].C1 != "2718824.476" {
		t.Errorf("C1 was incorrect, C1 was %s", geometries.Surfaces[0].Boundary[0].Polyline.Coords[0].C1)
	}
	if geometries.Surfaces[0].Boundary[0].Polyline.Coords[0].C2 != "1228990.278" {
		t.Errorf("C2 was incorrect, C1 was %s", geometries.Surfaces[0].Boundary[0].Polyline.Coords[0].C2)
	}
	if geometries.Surfaces[0].Boundary[0].Polyline.Coords[0].C3 != "" {
		t.Errorf("C3 was incorrect, C1 was %s", geometries.Surfaces[0].Boundary[0].Polyline.Coords[0].C3)
	}

	// Polygon (area)
	if geometries.Surfaces[2].Boundary[0].Polyline.Coords[0].C1 != "2718714.655" {
		t.Errorf("C1 was incorrect, C1 was %s", geometries.Surfaces[2].Boundary[0].Polyline.Coords[0].C1)
	}
	if geometries.Surfaces[2].Boundary[0].Polyline.Coords[0].C2 != "1228945.210" {
		t.Errorf("C2 was incorrect, C1 was %s", geometries.Surfaces[2].Boundary[0].Polyline.Coords[0].C2)
	}
	if geometries.Surfaces[2].Boundary[0].Polyline.Coords[0].C3 != "" {
		t.Errorf("C3 was incorrect, C1 was %s", geometries.Surfaces[2].Boundary[0].Polyline.Coords[0].C3)
	}

	// Polygon (surface with hole)
	if geometries.Surfaces[1].Boundary[0].Polyline.Coords[0].C1 != "2718709.475" {
		t.Errorf("C1 was incorrect, C1 was %s", geometries.Surfaces[1].Boundary[0].Polyline.Coords[0].C1)
	}
	if geometries.Surfaces[1].Boundary[1].Polyline.Coords[0].C1 != "2718734.168" {
		t.Errorf("C1 was incorrect, C1 was %s", geometries.Surfaces[1].Boundary[1].Polyline.Coords[0].C1)
	}
	if geometries.Surfaces[1].Boundary[0].Polyline.Coords[0].C2 != "1228935.713" {
		t.Errorf("C2 was incorrect, C1 was %s", geometries.Surfaces[1].Boundary[0].Polyline.Coords[0].C2)
	}
	if geometries.Surfaces[1].Boundary[1].Polyline.Coords[0].C2 != "1228930.533" {
		t.Errorf("C2 was incorrect, C1 was %s", geometries.Surfaces[1].Boundary[1].Polyline.Coords[0].C2)
	}
	if geometries.Surfaces[1].Boundary[0].Polyline.Coords[0].C3 != "" {
		t.Errorf("C3 was incorrect, C1 was %s", geometries.Surfaces[1].Boundary[0].Polyline.Coords[0].C3)
	}
	if geometries.Surfaces[1].Boundary[1].Polyline.Coords[0].C3 != "" {
		t.Errorf("C3 was incorrect, C1 was %s", geometries.Surfaces[1].Boundary[1].Polyline.Coords[0].C3)
	}
}
