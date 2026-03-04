package models

type Geometries interface {
	Polygons() [][]float64
	PolygonsWkt() []string
	Polygon() [][][][]float64
	PolygonWkt() []string
	Points() [][]float64
	PointsWkt() []string
	Point() [][]float64
	PointWkt() []string
	Lines() [][]float64
	LinesWkt() []string
	Line() [][][]float64
	LineWkt() []string
}
