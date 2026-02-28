package writer

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/tuquilla/interlis"
)

func TestGeojsonCreation(t *testing.T) {
	xtf, _ := os.Open("test.xtf")
	defer xtf.Close()
	decoder := xml.NewDecoder(xtf)

	geometries := interlis.ReadGeometry(decoder)
	Geojson(geometries)
}
