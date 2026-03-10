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

func TestPointsWkt(t *testing.T) {
	geometries := Geometries{MultiCoords: []MultiCoord{{Coords: []Coord{{C1: "1", C2: "2", C3: ""}, {C1: "6", C2: "8", C3: ""}, {C1: "9", C2: "22", C3: ""}}}}}
	wkt := geometries.PointsWkt()

	if wkt[0] != "MULTIPOINT ((1 2), (6 8), (9 22))" {
		t.Errorf("WKT MULTIPOINT is: %s", wkt[0])
	}
}

func TestLineWkt(t *testing.T) {
	geometries := Geometries{Polylines: []Polyline{{Coords: []Coord{{C1: "1", C2: "2", C3: ""}, {C1: "4", C2: "5", C3: ""}}}}}
	wkt := geometries.LineWkt()

	if wkt[0] != "LINESTRING (1 2, 4 5)" {
		t.Errorf("WKT LINESTRING is: %s", wkt[0])
	}
}

func TestLinesWKT(t *testing.T) {
	polyLine1 := Polyline{[]Coord{{C1: "-1", C2: "-1", C3: ""}, {C1: "2", C2: "2", C3: ""}}}
	polyLine2 := Polyline{[]Coord{{C1: "54", C2: "54", C3: ""}, {C1: "66", C2: "66", C3: ""}}}
	polyLine3 := Polyline{[]Coord{{C1: "3", C2: "3", C3: ""}, {C1: "6", C2: "6", C3: ""}, {C1: "9", C2: "9", C3: ""}}}
	multiPolyline := []Polyline{polyLine1, polyLine2, polyLine3}
	geometries := Geometries{MutliPolylines: []MultiPolyline{{multiPolyline}}}

	wkt := geometries.LinesWkt()

	if wkt[0] != "MULTILINESTRING ((-1 -1, 2 2), (54 54, 66 66), (3 3, 6 6, 9 9))" {
		t.Errorf("WKT MULTILINESTRING is: %s", wkt[0])
	}
}

func TestPolygonWkt(t *testing.T) {
	exteriorPolyline := Polyline{[]Coord{{C1: "0", C2: "0"}, {C1: "0", C2: "4"}, {C1: "4", C2: "4"}, {C1: "4", C2: "0"}, {C1: "0", C2: "0"}}}
	exterior := Exterior{Polyline: exteriorPolyline}
	interiorPolyline := Polyline{[]Coord{{C1: "1", C2: "1"}, {C1: "3", C2: "1"}, {C1: "3", C2: "3"}, {C1: "1", C2: "3"}, {C1: "1", C2: "1"}}}
	interior := Interior{Polyline: interiorPolyline}
	surface := Surface{Exterior: exterior, Interiors: []Interior{interior}}
	geometries := Geometries{Surfaces: []Surface{surface}}
	wkt := geometries.PolygonWkt()

	if wkt[0] != ("POLYGON ((0 0, 4 0, 4 4, 0 4, 0 0), (1 1, 1 3, 3 3, 3 1, 1 1))") {
		t.Errorf("WKT POLYGON is: %s", wkt[0])
	}

}

func TestPolygonsWkt(t *testing.T) {
	exteriorPolyline1 := Polyline{[]Coord{{C1: "0", C2: "0"}, {C1: "0", C2: "4"}, {C1: "4", C2: "4"}, {C1: "4", C2: "0"}, {C1: "0", C2: "0"}}}
	exteriorPolyline2 := Polyline{[]Coord{{C1: "10", C2: "10"}, {C1: "10", C2: "14"}, {C1: "14", C2: "14"}, {C1: "14", C2: "10"}, {C1: "10", C2: "10"}}}
	exteriorPolyline3 := Polyline{[]Coord{{C1: "100", C2: "100"}, {C1: "140", C2: "100"}, {C1: "140", C2: "140"}, {C1: "100", C2: "140"}, {C1: "100", C2: "100"}}}
	exterior1 := Exterior{Polyline: exteriorPolyline1}
	exterior2 := Exterior{Polyline: exteriorPolyline2}
	exterior3 := Exterior{Polyline: exteriorPolyline3}
	interiorPolyline := Polyline{[]Coord{{C1: "1", C2: "1"}, {C1: "3", C2: "1"}, {C1: "3", C2: "3"}, {C1: "1", C2: "3"}, {C1: "1", C2: "1"}}}
	interior := Interior{Polyline: interiorPolyline}
	surface1 := Surface{Exterior: exterior1, Interiors: []Interior{interior}}
	surface2 := Surface{Exterior: exterior2}
	surface3 := Surface{Exterior: exterior3}
	surfaces := []Surface{surface1, surface2, surface3}
	geometries := Geometries{MultiSurfaces: []MultiSurface{{surfaces}}}
	wkt := geometries.PolygonsWkt()

	if wkt[0] != ("MULTIPOLYGON (((0 0, 4 0, 4 4, 0 4, 0 0), (1 1, 1 3, 3 3, 3 1, 1 1)), ((10 10, 14 10, 14 14, 10 14, 10 10)), ((100 100, 140 100, 140 140, 100 140, 100 100)))") {
		t.Errorf("WKT MULTIPOLYGON is: %s", wkt[0])
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
