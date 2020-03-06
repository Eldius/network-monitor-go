package logger

import (
	"encoding/json"
	"log"
	"os"
)

/*
Debug prints an object to log as a JSON
*/
func Debug(obj interface{}) {
	f, err := os.OpenFile("/tmp/network-monitor.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	jsonBytes, _ := json.MarshalIndent(&obj, "", "  ")
	log.Println(string(jsonBytes))
}

/*
Println prints message to log
*/
func Println(msg string) {
	f, err := os.OpenFile("/tmp/network-monitor.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(msg)
}

/*
Println prints message to log
*/
func Printf(msg string, v ...interface{}) {
	f, err := os.OpenFile("/tmp/network-monitor.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Printf(msg, v)
}
