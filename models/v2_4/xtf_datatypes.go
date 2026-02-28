package v2_4

type Coord struct {
	C1 string `xml:"c1"`
	C2 string `xml:"c2"`
	C3 string `xml:"c3"`
}

type Polyline struct {
	Coords []Coord `xml:"coord"`
}

type Boundary struct {
	Polyline Polyline `xml:"polyline"`
}

type Surface struct {
	Boundary  Boundary   `xml:"boundary"`
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

func (geometries Geometries) Point() []float64 {
	return []float64{}
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
