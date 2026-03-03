package writer

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tuquilla/interlis/models"
	"github.com/tuquilla/interlis/models/geojson"
)

func Geojson(geometries models.Geometries, outputFilePath string) error {
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
		return nil
	}

	err = writeToFile(jsonResult, outputFilePath)
	if err != nil {
		return err
	}

	return nil
}

func writeToFile(output []byte, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("Error creating geojson file, error: %v", err)
	}
	defer f.Close()

	_, err = f.Write(output)
	if err != nil {
		return fmt.Errorf("Error writing to geojson file, error: %v", err)
	}
	return nil
}
