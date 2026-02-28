package writer

import (
	"encoding/json"
	"fmt"

	"github.com/tuquilla/interlis/models"
	"github.com/tuquilla/interlis/models/geojson"
)

func Geojson(geometries models.Geometries) {
	featureCollection := geojson.CreateFeatureCollection()

	// Point
	_, geometry := geometries.Point()
	feature := geojson.CreatePointFeature(geometry)
	featureCollection.AddFeature(feature)

	// MultiPoint

	// LineString

	//MultiLineString

	//Polygon

	//MultiPolygon

	jsonResult, err := json.Marshal(featureCollection)
	if err != nil {
		fmt.Println("Error marshalling json")
	}
	fmt.Println(string(jsonResult))
}
