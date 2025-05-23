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
		case "HL":
			// HL segment can be submitter, receiver, subscriber, or dependent
			// HL01, HL02 (optional), HL03, HL04
			if len(elements) < 4 || len(elements) > 5 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			hl := HL{
				HierarchicalIDNumber:  elements[1],
				HierarchicalLevelCode: elements[len(elements)-2],
				HierarchicalChildCode: elements[len(elements)-1],
			}
			if len(elements) == 5 {
				hl.HierarchicalParentID = elements[2]
			}
			// Assign HL to the correct place based on context or HL01
			if data.SubmitterHL.HierarchicalIDNumber == "" {
				data.SubmitterHL = hl
			} else if data.ReceiverHL.HierarchicalIDNumber == "" {
				data.ReceiverHL = hl
			} else if data.SubscriberHL.HierarchicalIDNumber == "" {
				data.SubscriberHL = hl
			} else {
				data.DependentHLs = append(data.DependentHLs, hl)
			}
		case "NM1":
			// NM1: 10 or 12 elements (some optional)
			if len(elements) < 4 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			nm1 := NM1{
				EntityIdentifierCode: elements[1],
				EntityTypeQualifier:  elements[2],
				LastNameOrOrgName:    elements[3],
			}
			if len(elements) > 4 {
				nm1.FirstName = elements[4]
			}
			if len(elements) > 5 {
				nm1.MiddleName = elements[5]
			}
			if len(elements) > 6 {
				nm1.NamePrefix = elements[6]
			}
			if len(elements) > 7 {
				nm1.NameSuffix = elements[7]
			}
			if len(elements) > 8 {
				nm1.IdentificationCodeQualifier = elements[8]
			}
			if len(elements) > 9 {
				nm1.IdentificationCode = elements[9]
			}
			if len(elements) > 10 {
				nm1.EntityRelationshipCode = elements[10]
			}
			if len(elements) > 11 {
				nm1.EntityIdentifierCode2 = elements[11]
			}
			// Attach NM1 to the last HL that doesn't have NM1
			if data.SubscriberHL.NM1.EntityIdentifierCode == "" {
				data.SubscriberHL.NM1 = nm1
			} else if data.ReceiverHL.NM1.EntityIdentifierCode == "" {
				data.ReceiverHL.NM1 = nm1
			} else if data.SubmitterHL.NM1.EntityIdentifierCode == "" {
				data.SubmitterHL.NM1 = nm1
			} else if len(data.DependentHLs) > 0 && data.DependentHLs[len(data.DependentHLs)-1].NM1.EntityIdentifierCode == "" {
				data.DependentHLs[len(data.DependentHLs)-1].NM1 = nm1
			}
		case "DMG":
			// DMG: up to 9 elements, all optional after DMG01
			if len(elements) < 2 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			dmg := &DMG{}
			if len(elements) > 1 {
				dmg.DateTimePeriodFormatQualifier = elements[1]
			}
			if len(elements) > 2 {
				dmg.DateOfBirth = elements[2]
			}
			if len(elements) > 3 {
				dmg.GenderCode = elements[3]
			}
			if len(elements) > 4 {
				dmg.MaritalStatusCode = elements[4]
			}
			if len(elements) > 5 {
				dmg.CitizenshipCode = elements[5]
			}
			if len(elements) > 6 {
				dmg.CountryCode = elements[6]
			}
			if len(elements) > 7 {
				dmg.BaselineResidenceCode = elements[7]
			}
			if len(elements) > 8 {
				dmg.Age = elements[8]
			}
			if len(elements) > 9 {
				dmg.AgeQualifier = elements[9]
			}
			// Attach DMG to the last HL that doesn't have DMG
			if data.SubscriberHL.DMG == nil {
				data.SubscriberHL.DMG = dmg
			} else if len(data.DependentHLs) > 0 && data.DependentHLs[len(data.DependentHLs)-1].DMG == nil {
				data.DependentHLs[len(data.DependentHLs)-1].DMG = dmg
			}
		case "HI":
			// HI: variable number of elements, each HIxx is 2 elements
			if len(elements) < 2 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			hi := HI{}
			for i := 1; i+1 < len(elements); i += 2 {
				hi.HealthCareCodeInfo = append(hi.HealthCareCodeInfo, HealthCareCodeInfo{
					CodeListQualifierCode: elements[i],
					IndustryCode:          elements[i+1],
				})
			}
			data.HI = append(data.HI, hi)
		case "UM":
			// UM: up to 8 elements, some optional
			if len(elements) < 4 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			um := UM{
				RequestCategoryCode:        elements[1],
				CertificationTypeCode:      elements[2],
				ServicesReviewDecisionCode: elements[3],
			}
			if len(elements) > 4 {
				um.MonetaryAmount = elements[4]
			}
			if len(elements) > 5 {
				um.Quantity = elements[5]
			}
			if len(elements) > 6 {
				um.RejectReasonCode = elements[6]
			}
			if len(elements) > 7 {
				um.ActionCode = elements[7]
			}
			if len(elements) > 8 {
				um.LevelOfServiceCode = elements[8]
			}
			data.UM = append(data.UM, um)
		case "BHT":
			if len(elements) < 6 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			data.BHT = BHT{
				HierarchicalStructureCode: elements[1],
				TransactionSetPurposeCode: elements[2],
				ReferenceIdentification:   elements[3],
				Date:                      elements[4],
				Time:                      elements[5],
			}
			if len(elements) > 6 {
				data.BHT.TransactionTypeCode = elements[6]
			}
		case "TRN":
			trn := TRN{
				TraceTypeCode:           elements[1],
				ReferenceIdentification: elements[2],
			}
			if len(elements) > 3 {
				trn.OriginatingCompanyID = elements[3]
			}
			if len(elements) > 4 {
				trn.ReferenceIdentification2 = elements[4]
			}
			data.TRN = append(data.TRN, trn)
		case "N3":
			n3 := N3{
				AddressInformation1: elements[1],
			}
			if len(elements) > 2 {
				n3.AddressInformation2 = elements[2]
			}
			data.N3 = append(data.N3, n3)
		case "N4":
			n4 := N4{
				CityName:        elements[1],
				StateOrProvince: elements[2],
				PostalCode:      elements[3],
			}
			data.N4 = append(data.N4, n4)
		case "PER":
			per := PER{
				ContactFunctionCode: elements[1],
			}
			if len(elements) > 2 {
				per.Name = elements[2]
			}
			if len(elements) > 3 {
				per.CommunicationNumbers = elements[3:]
			}
			data.PER = append(data.PER, per)
		case "DTP":
			if len(elements) < 4 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			dtp := DTP{
				DateTimeQualifier: elements[1],
				FormatQualifier:   elements[2],
				DateTimePeriod:    elements[3],
			}
			data.DTP = append(data.DTP, dtp)
		case "SV2":
			if len(elements) < 6 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			sv2 := SV2{
				ProductOrServiceID:            elements[1],
				ProcedureModifier:             "",
				LineItemChargeAmount:          elements[3],
				UnitOrBasisForMeasurementCode: elements[4],
				ServiceUnitCount:              elements[5],
			}
			if len(elements) > 2 {
				sv2.ProcedureModifier = elements[2]
			}
			data.SV2 = append(data.SV2, sv2)
		case "PWK":
			if len(elements) < 3 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			pwk := PWK{
				ReportTypeCode:         elements[1],
				ReportTransmissionCode: elements[2],
			}
			if len(elements) > 3 {
				pwk.ReportCopies = elements[3]
			}
			if len(elements) > 7 {
				pwk.ReportControlNumber = elements[7]
			}
			data.PWK = append(data.PWK, pwk)
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
		case "HSD":
			if len(elements) < 3 {
				log.Fatalf("Incorrect number of elements found in segment: %q", segment)
			}
			hsd := HSD{
				QuantityQualifier: elements[1],
				Quantity:          elements[2],
			}
			data.HSD = append(data.HSD, hsd)
		default:
			log.Fatalf("Skipping unknown segment type: %q\n", segment)
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
		log.Fatalf("Unable to write data to %s: %value\n", trasformedFilePath, err)
	}

	log.Println("Wrote data to output file at: ", trasformedFilePath)
}
