package v2_3

import (
	"encoding/xml"
	"fmt"

	"github.com/tuquilla/interlis/models/v2_4"
)

func ReadGeometry(decoder *xml.Decoder) v2_4.Geometries {
	var geometries v2_4.Geometries
	for {
		tok, err := decoder.Token()
		if err != nil {
			break
		}
		switch se := tok.(type) {
		case xml.StartElement:
			if se.Name.Local == "SURFACE" {
				var surface v2_4.Surface
				err := decoder.DecodeElement(&surface, &se)
				if err != nil {
					fmt.Println("Error at decoding SURFACE element")
					return geometries
				}
				geometries.Surfaces = append(geometries.Surfaces, surface)
			}
			if se.Name.Local == "POLYLINE" {
				var polyline v2_4.Polyline
				err := decoder.DecodeElement(&polyline, &se)
				if err != nil {
					fmt.Println("Error at decoding POLYLINE element")
					return geometries
				}
				geometries.Polylines = append(geometries.Polylines, polyline)
			}
			if se.Name.Local == "COORD" {
				var coord v2_4.Coord
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
