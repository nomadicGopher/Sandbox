package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// getInputData fetches field data from the input file and loads them into the data struct (representing JSON format).
func getInputData() (data Transaction278) {
	// TODO

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

	_, err = trasformedFile.WriteString("Hello world!")
	if err != nil {
		log.Fatalf("Unable to write data to %s: %v", trasformedFilePath, err)
	}

	log.Println("Wrote data to output file at: ", trasformedFilePath)
}
