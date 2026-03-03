package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"

	"github.com/tuquilla/interlis"
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
		if i == 1 {
			format = args[i]
		}

		if i == 2 {
			path, err := filepath.Abs(args[i])
			if err != nil {
				fmt.Printf("Error with input file path, error: %v", err)
				return
			}
			inputFilePath = path
		}

		if i == 3 {
			path, err := filepath.Abs(args[i])
			if err != nil {
				fmt.Printf("Error with input file path, error: %v", err)
				return
			}
			outputFilePath = path
		}
	}

	fmt.Println("Format: ", format)
	fmt.Println("Output Filepath: ", outputFilePath)

	xtf, _ := os.Open(inputFilePath)
	defer xtf.Close()
	decoder := xml.NewDecoder(xtf)
	geometries := interlis.ReadGeometry(decoder)

	switch format {
	case "geojson":
		err := writer.Geojson(geometries, outputFilePath)
		if err != nil {
			fmt.Println(err)
		}
	default:
		fmt.Printf("Format %s is not supported\n", format)
	}
}
