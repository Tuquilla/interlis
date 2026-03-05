# Interlis
Basic Library to consume Interlis-Files (.xtf) and extract geometries.
This is an early pre-alpha version!

## Usage
Use `ReadGeometry(decoder *xml.Decoder) models.Geometries` to read out geometries from xtf.

Use `Gpkg(geometries models.Geometries, outputFilePath string) error` to create gpkg from geometries

Use `Geojson(geometries models.Geometries, outputFilePath string) error` to create geojson from geometries

The binary can be used as client tool for the creation of the formats above.

**Arglist**:
1. format ("geojson", "gpkg")
2. path to interlis file
3. (optional) output path, if not delivered it is the same directory as the bin/exe, e.g. /home/user/test/testdata.geojson