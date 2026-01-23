package models

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGeojsonPoint(t *testing.T) {

	// prepare point
	pointCoordinate := []float64{2718659.49, 1228988.33}
	pointJsonGeometry := jsonPointGeometry{"Point", pointCoordinate}

	pointFeature := feature{"Feature", pointJsonGeometry}
	pointFeatureArr := []feature{pointFeature}
	pointFeatureCollection := featureCollection{"FeatureCollection", pointFeatureArr}

	jsonResult, err := json.Marshal(pointFeatureCollection)
	if err != nil {
		fmt.Println("Fehler")
	}

	if string(jsonResult) != "{\"type\":\"FeatureCollection\",\"features\":[{\"type\":\"Feature\",\"geometry\":{\"type\":\"Point\",\"coordinates\":[2718659.49,1228988.33]}}]}" {
		t.Errorf("GeoJson Point war nicht korrekt")
	}
}

func TestGeojsonLine(t *testing.T) {

	// prepare line
	pointCoordinates := make([][]float64, 0)
	pointCoordinate1 := []float64{2718659.49, 1228988.33}
	pointCoordinate2 := []float64{2718669.49, 1228998.33}
	pointCoordinates = append(pointCoordinates, pointCoordinate1, pointCoordinate2)
	pointJsonGeometry := jsonGeometry{"LineString", pointCoordinates}

	pointFeature := feature{"Feature", pointJsonGeometry}
	pointFeatureArr := []feature{pointFeature}
	pointFeatureCollection := featureCollection{"FeatureCollection", pointFeatureArr}

	jsonResult, err := json.Marshal(pointFeatureCollection)
	if err != nil {
		fmt.Println("Fehler")
	}

	if string(jsonResult) != "{\"type\":\"FeatureCollection\",\"features\":[{\"type\":\"Feature\",\"geometry\":{\"type\":\"LineString\",\"coordinates\":[[2718659.49,1228988.33],[2718669.49,1228998.33]]}}]}" {
		t.Errorf("GeoJson LineString war nicht korrekt")
	}
}
