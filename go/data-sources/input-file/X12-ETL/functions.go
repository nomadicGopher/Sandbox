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
		segment = strings.TrimSpace(segment)
		if segment == "" {
			continue
		}
		elements := strings.Split(segment, "*")
		if len(elements) == 0 || elements[0] == "" {
			continue
		}
		switch elements[0] {
		case "ISA":
			if len(elements) != 17 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			data.ISA = ISA{
				AuthorizationInfoQualifier:  elements[1],
				AuthorizationInfo:           elements[2],
				SecurityInfoQualifier:       elements[3],
				SecurityInfo:                elements[4],
				InterchangeIDQualifier1:     elements[5],
				InterchangeSenderID:         elements[6],
				InterchangeIDQualifier2:     elements[7],
				InterchangeReceiverID:       elements[8],
				InterchangeDate:             elements[9],
				InterchangeTime:             elements[10],
				InterchangeControlStandards: elements[11],
				InterchangeControlVersion:   elements[12],
				InterchangeControlNumber:    elements[13],
				AcknowledgmentRequested:     elements[14],
				UsageIndicator:              elements[15],
				ComponentElementSeparator:   elements[16],
			}
		case "GS":
			if len(elements) != 9 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			data.GS = GS{
				FunctionalIDCode:           elements[1],
				ApplicationSenderCode:      elements[2],
				ApplicationReceiverCode:    elements[3],
				Date:                       elements[4],
				Time:                       elements[5],
				GroupControlNumber:         elements[6],
				ResponsibleAgencyCode:      elements[7],
				VersionReleaseIndustryCode: elements[8],
			}
		case "ST":
			// ST03 is optional
			if len(elements) < 3 || len(elements) > 4 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			data.ST = ST{
				TransactionSetIDCode:        elements[1],
				TransactionSetControlNumber: elements[2],
			}
			if len(elements) == 4 {
				data.ST.ImplementationConventionRef = elements[3]
			}
		case "SE":
			if len(elements) != 3 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			// SE01 is int
			var numSegments int
			fmt.Sscanf(elements[1], "%d", &numSegments)
			data.SE = SE{
				NumberOfIncludedSegments:    numSegments,
				TransactionSetControlNumber: elements[2],
			}
		case "GE":
			if len(elements) != 3 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			data.GE = GE{
				NumberOfTransactionSetsIncluded: elements[1],
				GroupControlNumber:              elements[2],
			}
		case "IEA":
			if len(elements) != 3 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			data.IEA = IEA{
				NumberOfIncludedGroups:   elements[1],
				InterchangeControlNumber: elements[2],
			}
		default:
			log.Fatalf("Skipping unknown segment type: %q\n", segment)
		}
	}

	log.Println("Loaded input data into memory.")
	return data
}

// transformData makes a sample data transformation at the end of the data set.
func transformData(data Transaction278) (transformedData Transaction278) {
	transformedData = data

	transformedData.Body = "Body of requests goes here... This is a sample transformation."

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
		log.Fatalf("Unable to write data to %s: %value\n", trasformedFilePath, err)
	}

	log.Println("Wrote data to output file at: ", trasformedFilePath)
}
