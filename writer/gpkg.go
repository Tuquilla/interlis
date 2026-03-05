package writer

import (
	"fmt"

	gdal "github.com/lukeroth/gdal"
	"github.com/tuquilla/interlis/models"
)

func Gpkg(geometries models.Geometries, outputFilePath string) error {

	driver, err := gdal.GetDriverByName("GPKG")
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	dataset := driver.Create(outputFilePath, 0, 0, 0, gdal.Unknown, nil)
	defer dataset.Close()

	srs := gdal.CreateSpatialReference("")
	srs.FromEPSG(2056)

	var geometryList []string
	var layer gdal.Layer

	// Points
	geometryList = geometries.PointWkt()
	if geometryList != nil {
		layer = dataset.CreateLayer("points", srs, gdal.GT_Point, nil)
		for _, geometryElement := range geometryList {
			geometry, err := gdal.CreateFromWKT(geometryElement, srs)
			if err != nil {
				return fmt.Errorf("Error creating point layer for gpkg, err: %v", err)
			}

			feature := layer.Definition().Create()
			feature.SetGeometry(geometry)
			layer.Create(feature)
		}
	}

	// Lines
	geometryList = geometries.LineWkt()
	if geometryList != nil {
		layer = dataset.CreateLayer("lines", srs, gdal.GT_LineString, nil)
		for _, geometryElement := range geometryList {
			geometry, err := gdal.CreateFromWKT(geometryElement, srs)
			if err != nil {
				return fmt.Errorf("Error creating line layer for gpkg, err: %v", err)
			}

			feature := layer.Definition().Create()
			feature.SetGeometry(geometry)
			layer.Create(feature)
		}
	}

	// Polygons
	geometryList = geometries.PolygonWkt()
	if geometryList != nil {
		layer = dataset.CreateLayer("polygon", srs, gdal.GT_Polygon, nil)
		for _, geometryElement := range geometryList {
			geometry, err := gdal.CreateFromWKT(geometryElement, srs)
			if err != nil {
				return fmt.Errorf("Error creating polygon layer for gpkg, error: %v", err)
			}

			feature := layer.Definition().Create()
			feature.SetGeometry(geometry)
			layer.Create(feature)
		}
	}

	return nil
}
