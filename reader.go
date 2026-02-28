package interlis

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/tuquilla/interlis/models"
	"github.com/tuquilla/interlis/reader/v2_3"
	"github.com/tuquilla/interlis/reader/v2_4"
)

func ReadGeometry(decoder *xml.Decoder) models.Geometries {
	//var geometries models.Geometries
	var interlisVersion string
checkVersion:
	for {
		tok, err := decoder.Token()
		if err != nil {
			break
		}
		switch se := tok.(type) {
		case xml.StartElement:
			if strings.ToLower(se.Name.Local) == "transfer" {
				for _, element := range se.Attr {
					if element.Value == "http://www.interlis.ch/xtf/2.4/INTERLIS" {
						interlisVersion = "2.4"
						break checkVersion
					}
					if element.Value == "http://www.interlis.ch/INTERLIS2.3" {
						interlisVersion = "2.3"
						break checkVersion
					}
				}
				fmt.Println("No valid interlis version for this tool found")
				break checkVersion
			}
		}
	}
	// TODO Run Version based on interlisVersion
	var geometries models.Geometries
	if interlisVersion == "2.3" {
		geometries = v2_3.ReadGeometry(decoder)
	}

	if interlisVersion == "2.4" {
		geometries = v2_4.ReadGeometry(decoder)
	}
	return geometries
}
