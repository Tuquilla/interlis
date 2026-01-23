package models

import (
	"testing"
)

func TestReadGeometry(t *testing.T) {
	coord := Coord{"1", "2", "3"}
	geometryArray := coord.Geometry()
	if geometryArray[0] != "1" || geometryArray[1] != "2" || geometryArray[2] != "3" {
		t.Errorf("Geometry() doesn't return correct geometry array")
	}

	coord = Coord{"1", "2", ""}
	geometryArray = coord.Geometry()
	if len(geometryArray) != 2 {
		t.Errorf("Geometry() doesn't return correct geometry array if there is no C3 coordinate")
	}
}
