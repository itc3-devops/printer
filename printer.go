package printer

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Debug : Toggle for enabling / disable debug
var Debug bool

// LogLocation : Define location of local log file
var LogLocation string



func init () {
	// Configure logging

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Only log the warning severity or above.

	loglevel := os.Getenv("Debug")
	if loglevel == "true" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.ErrorLevel)
	}
	// Set log location

	LogLocation = os.Getenv("LogLocation")
	LogLocation = string(LogLocation)

	if LogLocation == "" {
		log.SetOutput(os.Stdout)
	} else {

		fileMgmt.CreateFile(LogLocation)
		//create your file with desired read/write permissions
		f, logErr := os.OpenFile(LogLocation, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if logErr != nil {
			fmt.Println("Print error: ", logErr)
		}
		//defer f.Close()
		log.SetOutput(f)
	}

}

// PrintDebug : Debug handling
func PrintDebug(label1, label2 message string) {
	// If Debug is set, print messages to screen
	if Debug == true {
		fmt.Println("Debug: " + label1, label2, message)
	}
		log.WithFields(log.Fields{label1: label2}).Debug(message)
}

// PrintInfo : Info handling
func PrintInfo(label1, label2, message string) {
	// If Debug is set, print messages to screen
	if Debug == true {
		fmt.Println("Info: " + label1, label2, message)
	}
		log.WithFields(log.Fields{label1: label2}).Info(message)
}

// PrintError : Error handling
func PrintError(err error, label1, label2, message string) {
	// If Debug is set, print messages to screen
	if Debug == true {
		fmt.Println("Error: " + label1, label2, message)
	}
		log.WithFields(log.Fields{label1: label2}).Error(message)
}


// PrintError : Error handling
func PrintFatal(err error, label1, label2, message string) {
	// If Debug is set, print messages to screen
	if Debug == true {
		fmt.Println("Fatal: " + label1, label2, message)
	}
		log.WithFields(log.Fields{label1: label2}).Fatal(message)
}