package writer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tuquilla/interlis/models"
	"github.com/tuquilla/interlis/models/geojson"
)

func Geojson(geometries models.Geometries, outputFilePath string) {
	featureCollection := geojson.CreateFeatureCollection()
	var feature geojson.Feature

	// Point
	pointGeometries := geometries.Point()
	for _, point := range pointGeometries {
		feature = geojson.CreatePointFeature(point)
		featureCollection.AddFeature(feature)
	}

	// MultiPoint
	multiPointGeometry := geometries.Points()
	if multiPointGeometry != nil {
		feature = geojson.CreateLineFeature("MultiPoint", multiPointGeometry)
		featureCollection.AddFeature(feature)
	}

	// LineString
	lineStrings := geometries.Line()
	for _, lineString := range lineStrings {
		feature = geojson.CreateLineFeature("LineString", lineString)
		featureCollection.AddFeature(feature)
	}

	//MultiLineString
	multiLineString := geometries.Lines()
	if multiLineString != nil {
		feature = geojson.CreateLineFeature("MultiLineString", multiLineString)
		featureCollection.AddFeature(feature)
	}

	//Polygon
	polygons := geometries.Polygon()
	for _, polygon := range polygons {
		feature = geojson.CreatePolygonFeature("Polygon", polygon)
		featureCollection.AddFeature(feature)
	}

	//MultiPolygon
	multiPolygon := geometries.Polygons()
	if multiPolygon != nil {
		feature = geojson.CreateLineFeature("MultiPolygon", multiPolygon)
		featureCollection.AddFeature(feature)
	}

	jsonResult, err := json.Marshal(featureCollection)
	if err != nil {
		fmt.Printf("Error marshalling json, %v", err)
	}

	if outputFilePath == "" {
		fmt.Println(string(jsonResult))
		return
	}

	writeToFile(jsonResult, outputFilePath)
	return
}

func writeToFile(output []byte, path string) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	f.Write(output)

}
