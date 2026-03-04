package v2_3

import (
	"fmt"
	"strconv"
	"strings"
)

type Coord struct {
	C1 string `xml:"C1"`
	C2 string `xml:"C2"`
	C3 string `xml:"C3"`
}

type Polyline struct {
	Coords []Coord `xml:"COORD"`
}

type Boundary struct {
	Polyline Polyline `xml:"POLYLINE"`
}

type Surface struct {
	Boundaries []Boundary `xml:"BOUNDARY"`
}

type Geometries struct {
	Surfaces  []Surface
	Polylines []Polyline
	Coords    []Coord
}

func (geometries Geometries) Point() [][]float64 {
	points := [][]float64(nil)

	for _, coord := range geometries.Coords {
		point := make([]float64, 2)

		coord1, err := strconv.ParseFloat(coord.C1, 64)
		if err != nil {
			fmt.Printf("Conversion of Point Coord failed, error: %s", err)
		}

		coord2, err := strconv.ParseFloat(coord.C2, 64)
		if err != nil {
			fmt.Printf("Conversion of Point Coord failed, error: %s", err)
		}

		point[0] = coord1
		point[1] = coord2
		points = append(points, point)
	}
	return points
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
	lines := [][][]float64(nil)

	for _, polyline := range geometries.Polylines {
		line := [][]float64(nil)
		for _, coord := range polyline.Coords {
			point := make([]float64, 2)
			coord1, err := strconv.ParseFloat(coord.C1, 64)
			if err != nil {
				fmt.Printf("Conversion of Point Coord failed, error: %s", err)
			}

			coord2, err := strconv.ParseFloat(coord.C2, 64)
			if err != nil {
				fmt.Printf("Conversion of Point Coord failed, error: %s", err)
			}
			point[0] = coord1
			point[1] = coord2
			line = append(line, point)
		}
		lines = append(lines, line)
	}

	return lines
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

func (geometries Geometries) Polygons() [][]float64 {
	return [][]float64(nil)
}

func (geometries Geometries) PolygonWkt() []string {
	polygons := []string(nil)

	for _, boundaries := range geometries.Surfaces {
		var polygon strings.Builder
		polygon.WriteString("POLYGON (")
		for _, boundary := range boundaries.Boundaries {
			polygon.WriteString("(")
			for index, coord := range boundary.Polyline.Coords {
				fmt.Fprintf(&polygon, "%s %s", coord.C1, coord.C2)
				if index < len(boundary.Polyline.Coords)-1 {
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

func (geometries Geometries) Polygon() [][][][]float64 {
	polygons := [][][][]float64(nil)

	for _, boundaries := range geometries.Surfaces {
		bound := [][][]float64(nil)
		for _, boundary := range boundaries.Boundaries {
			polygon := [][]float64(nil)
			for _, coord := range boundary.Polyline.Coords {
				point := make([]float64, 2)
				coord1, err := strconv.ParseFloat(coord.C1, 64)
				if err != nil {
					fmt.Printf("Conversion of Point Coord failed, error: %s", err)
				}

				coord2, err := strconv.ParseFloat(coord.C2, 64)
				if err != nil {
					fmt.Printf("Conversion of Point Coord failed, error: %s", err)
				}
				point[0] = coord1
				point[1] = coord2
				polygon = append(polygon, point)
			}
			bound = append(bound, polygon)
		}
		polygons = append(polygons, bound)
	}

	return polygons
}

func (geometries Geometries) PolygonsWkt() []string { return nil }

func (boundary *Boundary) isPolygonClockwise() bool {
	sum := 0.0
	for index := 0; index < len(boundary.Polyline.Coords)-1; index++ {
		coords := boundary.Polyline.Coords
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

// Todo implement
func (Boundary *Boundary) inversePolygonOrientation() {

}
