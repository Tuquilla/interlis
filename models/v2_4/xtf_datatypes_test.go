package v2_4

import (
	"testing"
)

func TestPointWkt(t *testing.T) {
	geometries := Geometries{Coords: []Coord{{C1: "1", C2: "2", C3: ""}}}
	wkt := geometries.PointWkt()

	if wkt[0] != "POINT (1 2)" {
		t.Errorf("WKT POINT is: %s", wkt[0])
	}
}

func TestLineWkt(t *testing.T) {
	geometries := Geometries{Polylines: []Polyline{{Coords: []Coord{{C1: "1", C2: "2", C3: ""}, {C1: "4", C2: "5", C3: ""}}}}}
	wkt := geometries.LineWkt()

	if wkt[0] != "LINESTRING (1 2, 4 5)" {
		t.Errorf("WKT LINESTRING is: %s", wkt[0])
	}
}

func TestPolyongWkt(t *testing.T) {
	exteriorPolyline := Polyline{[]Coord{{C1: "0", C2: "0"}, {C1: "0", C2: "4"}, {C1: "4", C2: "4"}, {C1: "4", C2: "0"}, {C1: "0", C2: "0"}}}
	exterior := Exterior{Polyline: exteriorPolyline}
	interiorPolyline := Polyline{[]Coord{{C1: "1", C2: "1"}, {C1: "3", C2: "1"}, {C1: "3", C2: "3"}, {C1: "1", C2: "3"}, {C1: "1", C2: "1"}}}
	interior := Interior{Polyline: interiorPolyline}
	surface := Surface{Exterior: exterior, Interiors: []Interior{interior}}
	geometries := Geometries{Surfaces: []Surface{surface}}
	wkt := geometries.PolygonWkt()

	if wkt[0] != ("POLYGON ((0 0, 4 0, 4 4, 0 4, 0 0), (1 1, 1 3, 3 3, 3 1, 1 1))") {
		t.Errorf("WKT Polygon is: %s", wkt[0])
	}

}

func TestIsClockwise(t *testing.T) {
	clockwiseCoords := []Coord{{C1: "0", C2: "0"}, {C1: "0", C2: "1"}, {C1: "1", C2: "1"}, {C1: "1", C2: "0"}, {C1: "0", C2: "0"}}
	polyline := Polyline{Coords: clockwiseCoords}

	isClockwise := polyline.IsPolygonClockwise()

	if isClockwise == false {
		t.Errorf("Polygon should be clockwise but isn't")
	}
}

func TestInversePolygonOrientation(t *testing.T) {
	counterclockwiseCoords := []Coord{{C1: "0", C2: "0"}, {C1: "1", C2: "0"}, {C1: "1", C2: "1"}, {C1: "0", C2: "1"}, {C1: "0", C2: "0"}}
	polyline := Polyline{Coords: counterclockwiseCoords}
	isClockwise := polyline.IsPolygonClockwise()

	if isClockwise == true {
		t.Errorf("Polygon should be counterclockwise but isn't")
	}

	polyline.InversePolygonOrientation()
	isClockwise = polyline.IsPolygonClockwise()

	if isClockwise == false {
		t.Errorf("Inversed polygon should be clockwise but isn't")
	}

}
