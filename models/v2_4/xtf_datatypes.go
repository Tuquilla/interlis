package v2_4

import (
	"fmt"
	"strconv"
	"strings"
)

type Coord struct {
	C1 string `xml:"c1"`
	C2 string `xml:"c2"`
	C3 string `xml:"c3"`
}

type Polyline struct {
	Coords []Coord `xml:"coord"`
}

type Surface struct {
	Exterior  Exterior   `xml:"exterior"`
	Interiors []Interior `xml:"interior"`
}

type MultiSurface struct {
	Surfaces []Surface `xml:"surface"`
}

type Exterior struct {
	Polyline Polyline `xml:"polyline"`
}

type Interior struct {
	Polyline Polyline `xml:"polyline"`
}

type Geometries struct {
	Surfaces      []Surface
	MultiSurfaces []MultiSurface
	Polylines     []Polyline
	Coords        []Coord
}

func (geometries Geometries) Point() [][]float64 {
	return [][]float64(nil)
}

func (geometries Geometries) PointWkt() []string {
	points := []string(nil)

	for _, coord := range geometries.Coords {
		points = append(points, fmt.Sprintf("POINT (%s %s)", coord.C1, coord.C2))
	}

	return points
}

func (geometries Geometries) Points() [][]float64 {
	return [][]float64(nil)
}

func (geometries Geometries) PointsWkt() []string { return nil }

func (geometries Geometries) Line() [][][]float64 {
	return [][][]float64(nil)
}

func (geometries Geometries) LineWkt() []string {
	lines := []string(nil)

	for _, polyline := range geometries.Polylines {
		var line strings.Builder
		line.WriteString("LINESTRING (")
		for index, coord := range polyline.Coords {
			fmt.Fprintf(&line, "%s %s", coord.C1, coord.C2)
			if index < len(polyline.Coords)-1 {
				line.WriteString(", ")
			}
		}
		line.WriteString(")")
		lines = append(lines, line.String())
	}

	return lines
}

func (geometries Geometries) Lines() [][]float64 {
	return [][]float64(nil)
}

func (geometries Geometries) LinesWkt() []string { return nil }

func (geometries Geometries) Polygon() [][][][]float64 {
	return [][][][]float64(nil)
}

func (geometries Geometries) PolygonWkt() []string {
	polygons := []string(nil)

	for _, surface := range geometries.Surfaces {
		var polygon strings.Builder
		polygon.WriteString("POLYGON (")

		if surface.Exterior.Polyline.IsPolygonClockwise() == true {
			surface.Exterior.Polyline.InversePolygonOrientation()
		}

		polygon.WriteString("(")
		for index, coord := range surface.Exterior.Polyline.Coords {
			fmt.Fprintf(&polygon, "%s %s", coord.C1, coord.C2)
			if index < len(surface.Exterior.Polyline.Coords)-1 {
				polygon.WriteString(", ")
			}
		}
		polygon.WriteString(")")

		for interiorIndex := range surface.Interiors {
			polygon.WriteString(", ")
			if surface.Interiors[interiorIndex].Polyline.IsPolygonClockwise() == false {
				surface.Interiors[interiorIndex].Polyline.InversePolygonOrientation()
			}
			polygon.WriteString("(")
			for index, coord := range surface.Interiors[interiorIndex].Polyline.Coords {
				fmt.Fprintf(&polygon, "%s %s", coord.C1, coord.C2)
				if index < len(surface.Interiors[interiorIndex].Polyline.Coords)-1 {
					polygon.WriteString(", ")
				}
			}
			polygon.WriteString(")")
		}
		polygon.WriteString(")")
		polygons = append(polygons, polygon.String())
	}

	return polygons
}

func (geometries Geometries) Polygons() [][]float64 {
	return [][]float64(nil)
}

func (geometries Geometries) PolygonsWkt() []string { return nil }

func (polyline *Polyline) IsPolygonClockwise() bool {
	sum := 0.0
	coords := polyline.Coords
	for index := 0; index < len(coords)-1; index++ {
		x1, _ := strconv.ParseFloat(coords[index].C1, 64)
		x2, _ := strconv.ParseFloat(coords[index+1].C1, 64)
		y1, _ := strconv.ParseFloat(coords[index].C2, 64)
		y2, _ := strconv.ParseFloat(coords[index+1].C2, 64)
		sum += ((x2 - x1) * (y2 + y1))
	}
	sum = 0.5 * sum

	if sum > 0 {
		return true
	}
	return false
}

func (polyline *Polyline) InversePolygonOrientation() {
	coords := polyline.Coords
	for index := 0; index < len(coords)/2; index++ {
		firstCoord := coords[index]
		secondCoord := coords[len(coords)-(index+1)]
		coords[index] = secondCoord
		coords[len(coords)-(index+1)] = firstCoord
	}
}
