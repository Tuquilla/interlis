package v2_3

import (
	"fmt"
	"strconv"
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

func (geometries Geometries) Points() [][]float64 {
	return [][]float64(nil)
}

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

func (geometries Geometries) Lines() [][]float64 {
	return [][]float64(nil)
}

func (geometries Geometries) Polygons() [][]float64 {
	return [][]float64(nil)
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
