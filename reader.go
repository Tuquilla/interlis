package interlis

import (
	"encoding/xml"
	"fmt"

	"github.com/tuquilla/interlis/models"
)

func ReadGeometry(decoder *xml.Decoder) models.Geometries {
	var geometries models.Geometries
	for {
		tok, err := decoder.Token()
		if err != nil {
			break
		}
		switch se := tok.(type) {
		case xml.StartElement:
			if se.Name.Local == "SURFACE" {
				var surface models.Surface
				err := decoder.DecodeElement(&surface, &se)
				if err != nil {
					fmt.Println("Error at decoding SURFACE element")
					return geometries
				}
				geometries.Surfaces = append(geometries.Surfaces, surface)
			}
			if se.Name.Local == "POLYLINE" {
				var polyline models.Polyline
				err := decoder.DecodeElement(&polyline, &se)
				if err != nil {
					fmt.Println("Error at decoding POLYLINE element")
					return geometries
				}
				geometries.Polylines = append(geometries.Polylines, polyline)
			}
			if se.Name.Local == "COORD" {
				var coord models.Coord
				err := decoder.DecodeElement(&coord, &se)
				if err != nil {
					fmt.Println("Error at decoding COORD element")
					return geometries
				}
				geometries.Coords = append(geometries.Coords, coord)
			}
		}
	}
	return geometries
}
