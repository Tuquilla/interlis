package models

type Geometries interface {
	Polygons() (string, [][]float64)
	Polygon() (string, []float64)
	Points() (string, [][]float64)
	Point() (string, []float64)
	Lines() (string, [][]float64)
	Line() (string, []float64)
}
