package geojson

type JsonPointGeometry struct {
	GeojsonType string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type JsonLineGeometry struct {
	GeojsonType string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

type JsonPolygonGeometry struct {
	GeojsonType string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

type jsonGeometryInterface interface {
}

type Feature struct {
	FeatureType string                `json:"type"`
	Geometry    jsonGeometryInterface `json:"geometry"`
}

type FeatureCollection struct {
	FeatureCollectionType string    `json:"type"`
	Features              []Feature `json:"features"`
}

func CreateFeatureCollection() *FeatureCollection {
	return &FeatureCollection{FeatureCollectionType: "FeatureCollection", Features: []Feature{}}
}

func (fc *FeatureCollection) AddFeature(feature Feature) {
	fc.Features = append(fc.Features, feature)
}

//func CreateJsonGeometry(geojsonType string, geometry [][]float64) JsonGeometry {
//	return JsonGeometry{GeojsonType: geojsonType}
//}

func CreatePointFeature(point []float64) Feature {
	return Feature{FeatureType: "Feature", Geometry: JsonPointGeometry{GeojsonType: "Point", Coordinates: point}}
}

func CreateLineFeature(geometryType string, geometry [][]float64) Feature {
	return Feature{FeatureType: "Feature", Geometry: JsonLineGeometry{GeojsonType: geometryType, Coordinates: geometry}}
}

func CreatePolygonFeature(geometryType string, geometry [][][]float64) Feature {
	return Feature{FeatureType: "Feature", Geometry: JsonPolygonGeometry{GeojsonType: geometryType, Coordinates: geometry}}
}
