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

type geometry interface {
	Geometry() []string
}

func (coord *Coord) Geometry() []string {
	geometryArr := make([]string, 0)

	geometryArr = append(geometryArr, coord.C1)
	geometryArr = append(geometryArr, coord.C2)
	if coord.C3 != "" {
		geometryArr = append(geometryArr, coord.C3)
	}

	return geometryArr
}
