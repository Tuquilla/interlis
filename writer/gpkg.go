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

	// Points
	geometryList = geometries.PointWkt()
	if geometryList != nil {
		err := createLayer(&dataset, srs, "points", geometryList, gdal.GT_Point)
		if err != nil {
			return err
		}
	}

	// Multipoints
	geometryList = geometries.PointsWkt()
	if geometryList != nil {
		err := createLayer(&dataset, srs, "multipoints", geometryList, gdal.GT_MultiPoint)
		if err != nil {
			return err
		}
	}

	// Lines
	geometryList = geometries.LineWkt()
	if geometryList != nil {
		err := createLayer(&dataset, srs, "lines", geometryList, gdal.GT_LineString)
		if err != nil {
			return err
		}
	}

	// Multilines
	geometryList = geometries.LinesWkt()
	if geometryList != nil {
		err := createLayer(&dataset, srs, "multilines", geometryList, gdal.GT_MultiLineString)
		if err != nil {
			return err
		}
	}

	// Polygons
	geometryList = geometries.PolygonWkt()
	if geometryList != nil {
		err := createLayer(&dataset, srs, "polygons", geometryList, gdal.GT_Polygon)
		if err != nil {
			return err
		}
	}

	// Multipolygons
	geometryList = geometries.PolygonsWkt()
	if geometryList != nil {
		err := createLayer(&dataset, srs, "multipolygons", geometryList, gdal.GT_MultiPolygon)
		if err != nil {
			return err
		}
	}

	return nil
}

func createLayer(dataset *gdal.Dataset, srs gdal.SpatialReference, layername string, geometryList []string, geometryTyp gdal.GeometryType) error {
	layer := dataset.CreateLayer(layername, srs, geometryTyp, nil)
	for _, geometryElement := range geometryList {
		geometry, err := gdal.CreateFromWKT(geometryElement, srs)
		if err != nil {
			return fmt.Errorf("Error creating layer: %s, error: %v", layername, err)
		}

		feature := layer.Definition().Create()
		feature.SetGeometry(geometry)
		layer.Create(feature)
	}
	return nil
}
