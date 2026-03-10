package v2_4

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
			switch se.Name.Local {
			case "multisurface":
				var multisurface v2_4.MultiSurface
				err := decoder.DecodeElement(&multisurface, &se)
				if err != nil {
					fmt.Println("Error at decoding SURFACE element")
					return geometries
				}
				geometries.MultiSurfaces = append(geometries.MultiSurfaces, multisurface)
			case "surface":
				var surface v2_4.Surface
				err := decoder.DecodeElement(&surface, &se)
				if err != nil {
					fmt.Println("Error at decoding SURFACE element")
					return geometries
				}
				geometries.Surfaces = append(geometries.Surfaces, surface)
			case "multipolyline":
				var multipolyline v2_4.MultiPolyline
				err := decoder.DecodeElement(&multipolyline, &se)
				if err != nil {
					fmt.Println("Error at decoding POLYLINE element")
					return geometries
				}
				geometries.MutliPolylines = append(geometries.MutliPolylines, multipolyline)
			case "polyline":
				var polyline v2_4.Polyline
				err := decoder.DecodeElement(&polyline, &se)
				if err != nil {
					fmt.Println("Error at decoding POLYLINE element")
					return geometries
				}
				geometries.Polylines = append(geometries.Polylines, polyline)
			case "multicoord":
				var multiCoord v2_4.MultiCoord
				err := decoder.DecodeElement(&multiCoord, &se)
				if err != nil {
					fmt.Println("Error at decoding COORD element")
					return geometries
				}
				geometries.MultiCoords = append(geometries.MultiCoords, multiCoord)
			case "coord":
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
