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
	inFilePath := flag.String("inPath", "", "Full path to input file.")
	produceLogFile := flag.Bool("log", false, "Set to true if a log file should be produced.")
	flag.Parse()

	baseFileName := filepath.Base(*inFilePath)
	baseFileName = baseFileName[:len(baseFileName)-len(filepath.Ext(*inFilePath))]
	timeStamp := time.Now()
	formattedTimeStamp := timeStamp.Format("_2006-01-02_15-04-05")

	if *produceLogFile {
		logFilePath := filepath.Join(filepath.Dir(*inFilePath), baseFileName+formattedTimeStamp+".log")
		logFile, err := os.Create(logFilePath)
		if err != nil {
			log.Fatalln("Unable to create Log File: ", err)
		}
		defer logFile.Close()
		log.Println("Log file created at: ", logFilePath)
		multiWriter := io.MultiWriter(logFile, os.Stdout)
		log.SetOutput(multiWriter)
	}

	data, err := getInputData(*inFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	transformedData, err := transformData(data)
	if err != nil {
		log.Fatalln(err)
	}

	if err := writeTransformedData(*inFilePath, baseFileName, formattedTimeStamp, transformedData); err != nil {
		log.Fatalln(err)
	}

	os.Exit(0)
}
