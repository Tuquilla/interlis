package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	interlis "github.com/tuquilla/interlis/reader"
	"github.com/tuquilla/interlis/writer"
)

// using lib as cli tool
func main() {
	args := os.Args
	var format string
	var inputFilePath string
	var outputFilePath string

	for i := 1; i < len(args); i++ {

		// Check options (needs --option)

		switch i {
		case 1:
			format = args[i]

		case 2:
			path, err := filepath.Abs(args[i])
			if err != nil {
				fmt.Printf("Error with input file path, error: %v", err)
				return
			}
			inputFilePath = path

		case 3:
			path, err := filepath.Abs(args[i])
			if err != nil {
				fmt.Printf("Error with input file path, error: %v", err)
				return
			}
			outputFilePath = path
		}
	}

	fmt.Println("Format:", format)
	fmt.Println("Output Filepath:", outputFilePath)

	xtf, _ := os.Open(inputFilePath)
	defer xtf.Close()
	decoder := xml.NewDecoder(xtf)
	geometries := interlis.ReadGeometry(decoder)

	switch format {
	case "geojson":
		err := writer.Geojson(geometries, outputFilePath)
		if err != nil {
			fmt.Println(err.Error())
		}
	case "gpkg":
		err := writer.Gpkg(geometries, outputFilePath)
		if err != nil {
			fmt.Println(err.Error())
		}
	default:
		fmt.Printf("Format %s is not supported\n", format)
	}
}
