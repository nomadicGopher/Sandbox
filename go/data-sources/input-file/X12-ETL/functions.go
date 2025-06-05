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
func getInputData(inFilePath string) (data X12_278, err error) {
	inFileData, err := os.ReadFile(inFilePath)
	if err != nil {
		return data, fmt.Errorf("unable to open input file: %w", err)
	}
	segments := strings.Split(string(inFileData), "~")

	var currentGroup *Group

	for _, segment := range segments {
		segment = strings.TrimSpace(segment)

		if segment == "" {
			continue
		}

		elements := strings.Split(segment, "*")

		for i := range elements {
			elements[i] = strings.TrimSpace(elements[i])
		}

		switch elements[0] {
		case "ISA":
			if len(elements) != 17 {
				return data, fmt.Errorf("incorrect number of elements found in segment: %q", segment)
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
				return data, fmt.Errorf("incorrect number of elements found in segment: %q", segment)
			}

			currentGroup = &Group{
				GS: GS{
					FunctionalIDCode:           elements[1],
					ApplicationSenderCode:      elements[2],
					ApplicationReceiverCode:    elements[3],
					Date:                       elements[4],
					Time:                       elements[5],
					GroupControlNumber:         elements[6],
					ResponsibleAgencyCode:      elements[7],
					VersionReleaseIndustryCode: elements[8],
				},
			}
		case "GE":
			if len(elements) != 3 {
				return data, fmt.Errorf("incorrect number of elements found in segment: %q", segment)
			}

			if currentGroup == nil {
				return data, fmt.Errorf("GE segment found without matching GS segment: %q", segment)
			}

			currentGroup.GE = GE{
				NumberOfTransactionSetsIncluded: elements[1],
				GroupControlNumber:              elements[2],
			}

			data.GroupBatches = append(data.GroupBatches, *currentGroup)
		case "IEA":
			if len(elements) != 3 {
				return data, fmt.Errorf("incorrect number of elements found in segment: %q", segment)
			}

			data.IEA = IEA{
				NumberOfIncludedGroups:   elements[1],
				InterchangeControlNumber: elements[2],
			}
		default:
			return data, fmt.Errorf("found unknown segment type: %q", segment)
		}
	}

	log.Println("Loaded input data into memory.")
	return data, nil
}

// transformData makes a sample data transformation at the end of the data set.
func transformData(data X12_278) (transformedData X12_278, err error) {
	transformedData = data

	for i := range transformedData.GroupBatches {
		transformedData.GroupBatches[i].Body = fmt.Sprintf("GroupBatch[%d]: Transaction details should begin with an ST segment and end with an SE segment...", i)
		log.Printf("Updated transaction body for group[%d].\n", i)
	}

	return transformedData, nil
}

// writeTransformedData writes the transformed data into the transformed JSON file before being imported into a system.
func writeTransformedData(inFilePath, baseFileName, formattedTimeStamp string, transformedData X12_278) error {
	trasformedFilePath := filepath.Join(filepath.Dir(inFilePath), baseFileName+"_transformed"+formattedTimeStamp+".json")
	trasformedFile, err := os.Create(trasformedFilePath)
	if err != nil {
		return fmt.Errorf("unable to create Transformed File: %w", err)
	}
	defer trasformedFile.Close()

	transformedDataBytes, err := json.MarshalIndent(transformedData, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal data to JSON: %w", err)
	}

	_, err = trasformedFile.Write(transformedDataBytes)
	if err != nil {
		return fmt.Errorf("unable to write data to %s: %v", trasformedFilePath, err)
	}

	log.Println("Wrote data to output file at: ", trasformedFilePath)
	return nil
}
