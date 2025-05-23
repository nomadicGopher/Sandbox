package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// getInputData fetches field data from the input file and loads them into the data struct (representing JSON format).
func getInputData(inFilePath string) (data Transaction278) {
	inFileData, err := os.ReadFile(inFilePath)
	if err != nil {
		log.Fatalf("Unable to open input file: %v\n", err)
	}
	segments := strings.Split(string(inFileData), "~")
	for _, segment := range segments {
		elements := strings.Split(segment, "*")

		switch elements[0] {
		case "ISA":
			for i, element := range elements {
				log.Printf("Segment ISA%02d: %s", i+1, element)
				// TODO: load into struct field via variadic function?
			}
		default:
		}
	}

	log.Println("Loaded input data into memory.")
	return data
}

// transformData makes a sample data transformation at the end of the data set.
func transformData(data Transaction278, formattedTimeStamp string) (transformedData Transaction278) {
	transformedData = data

	transformedData.Comment = fmt.Sprint("Data transformed at: ", formattedTimeStamp)

	log.Println("Transformed data.")
	return transformedData
}

// writeTransformedData writes the transformed data into the transformed JSON file before being imported into a system.
func writeTransformedData(inFilePath, baseFileName, formattedTimeStamp string, transformedData Transaction278) {
	trasformedFilePath := filepath.Join(filepath.Dir(inFilePath), baseFileName+"_transformed"+formattedTimeStamp+".json")
	trasformedFile, err := os.Create(trasformedFilePath)
	if err != nil {
		log.Fatalln("Unable to create Transformed File: ", err)
	}
	defer trasformedFile.Close()

	transformedDataBytes, err := json.MarshalIndent(transformedData, "", "  ")
	if err != nil {
		log.Fatalln("Unable to marshal data to JSON: ", err)
	}

	_, err = trasformedFile.Write(transformedDataBytes)
	if err != nil {
		log.Fatalf("Unable to write data to %s: %v\n", trasformedFilePath, err)
	}

	log.Println("Wrote data to output file at: ", trasformedFilePath)
}
