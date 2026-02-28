package geojson

type JsonPointGeometry struct {
	GeojsonType string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type JsonGeometry struct {
	GeojsonType string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
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

func CreateFeatureCollection(feature []Feature) FeatureCollection {
	return FeatureCollection{FeatureCollectionType: "FeatureCollection", Features: feature}
}

func CreateJsonGeometry(geojsonType string) JsonGeometry {
	return JsonGeometry{GeojsonType: geojsonType}
}

func CreateJsonPointGeometry(point []float64) JsonPointGeometry {
	return JsonPointGeometry{GeojsonType: "Point", Coordinates: point}
}
