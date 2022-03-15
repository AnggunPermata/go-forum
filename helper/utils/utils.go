package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var Config = Configuration{}
var Logger *log.Logger
type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeOut int64
	Static       string
}

func P(a ...interface{}) {
	fmt.Println(a)
}

func InitLogger() {
	file, err := os.OpenFile("go-forum.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("failed to open log file - ", err)
	}
	Logger = log.New(file, "INFO", log.Ldate|log.Ltime|log.Lshortfile)
}

// a function to help load data from config.json
func LoadConfig() {
	file, err := os.Open("config.json")
	if err != nil {
		log.Fatalln("cannot open config.json file - ", err)
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&Config); err != nil {
		log.Fatalln("cannot get configuration from file - ", err)
	}
}

