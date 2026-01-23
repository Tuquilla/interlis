package models

type jsonPointGeometry struct {
	GeojsonType string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type jsonGeometry struct {
	GeojsonType string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

type jsonGeometryInterface interface {
}

type feature struct {
	FeatureType string                `json:"type"`
	Geometry    jsonGeometryInterface `json:"geometry"`
}

type featureCollection struct {
	FeatureCollectionType string    `json:"type"`
	Features              []feature `json:"features"`
}
