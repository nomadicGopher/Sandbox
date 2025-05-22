/*
	Companion Guide: https://www.cms.gov/files/document/esmd-x12n-278-companion-guide.pdf
		- ISA/GS segments (interchange and functional group headers)
		- ST/SE segments (transaction set headers and trailers)
		- HL loops for the submitter, receiver, subscriber (employee), and dependents
		- NM1 segments for names
		- DMG segments for demographics
		- HI segments for service information
		- UM segments for utilization management
*/

package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	inFilePath := flag.String("inPath", "./sample278.txt", "Full path to input file.")
	flag.Parse()
	timeStamp := time.Now()
	formattedTimeStamp := timeStamp.Format("_2006-01-02_15-04-05")

	baseFileName := filepath.Base(*inFilePath)
	baseFileName = baseFileName[:len(baseFileName)-len(filepath.Ext(*inFilePath))]

	logFilePath := filepath.Join(filepath.Dir(*inFilePath), baseFileName+formattedTimeStamp+".log")
	logFile, err := os.Create(logFilePath)
	if err != nil {
		log.Println("Unable to create Log File: ", err)
	}
	defer logFile.Close()

	trasformedFilePath := filepath.Join(filepath.Dir(*inFilePath), baseFileName+"_transformed"+formattedTimeStamp+filepath.Ext(*inFilePath))
	trasformedFile, err := os.Create(trasformedFilePath)
	if err != nil {
		log.Println("Unable to create Transformed File: ", err)
	}
	defer trasformedFile.Close()

	_, err = trasformedFile.WriteString("Hello world!")
	if err != nil {
		log.Println("Unable to write to Transformed File: ", err)
	}

	os.Exit(0)
}
