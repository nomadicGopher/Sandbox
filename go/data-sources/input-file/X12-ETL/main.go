package main

import (
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	inFilePath := flag.String("inPath", ".vscode/hhpcr_x12n278.txt", "Full path to input file.")
	flag.Parse()
	timeStamp := time.Now()
	formattedTimeStamp := timeStamp.Format("_2006-01-02_15-04-05")

	baseFileName := filepath.Base(*inFilePath)
	baseFileName = baseFileName[:len(baseFileName)-len(filepath.Ext(*inFilePath))]

	logFilePath := filepath.Join(filepath.Dir(*inFilePath), baseFileName+formattedTimeStamp+".log")
	logFile, err := os.Create(logFilePath)
	if err != nil {
		log.Fatalln("Unable to create Log File: ", err)
	}
	defer logFile.Close()
	log.Println("Log file created at: ", logFilePath)
	multiWriter := io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(multiWriter)

	data := getInputData(*inFilePath)

	transformedData := transformData(data, formattedTimeStamp)

	writeTransformedData(*inFilePath, baseFileName, formattedTimeStamp, transformedData)

	os.Exit(0)
}
