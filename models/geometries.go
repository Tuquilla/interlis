package models

type Geometries interface {
	Polygons() [][]float64
	Polygon() [][][][]float64
	Points() [][]float64
	Point() [][]float64
	Lines() [][]float64
	Line() [][][]float64
}
