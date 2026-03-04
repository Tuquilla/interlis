package writer

import (
	"encoding/xml"
	"os"
	"testing"

	reader "github.com/tuquilla/interlis/reader"
)

func TestGPKGCreation(t *testing.T) {
	xtf, _ := os.Open("test.xtf")
	defer xtf.Close()
	decoder := xml.NewDecoder(xtf)

	geometries := reader.ReadGeometry(decoder)
	Gpkg(geometries, "test.gpkg")
}
