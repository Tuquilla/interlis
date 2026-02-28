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
	Boundary []Boundary `xml:"BOUNDARY"`
}

type Geometries struct {
	Surfaces  []Surface
	Polylines []Polyline
	Coords    []Coord
}

func (geometries Geometries) Point() []float64 {
	point := make([]float64, 2)
	for _, coords := range geometries.Coords {
		coord1, err := strconv.ParseFloat(coords.C1, 64)
		if err != nil {
			fmt.Printf("Conversion of Point Coord failed, error: %s", err)
		}
		coord2, err := strconv.ParseFloat(coords.C2, 64)
		if err != nil {
			fmt.Printf("Conversion of Point Coord failed, error: %s", err)
		}
		point = append(point, coord1, coord2)
	}
	return point
}

func (geometries Geometries) Points() [][]float64 {
	return [][]float64{}
}

func (geometries Geometries) Line() []float64 {
	return []float64{}
}

func (geometries Geometries) Lines() [][]float64 {
	return [][]float64{}
}

func (geometries Geometries) Polygon() []float64 {
	return []float64{}
}

func (geometries Geometries) Polygons() [][]float64 {
	return [][]float64{}
}
