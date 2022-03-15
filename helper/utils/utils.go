package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var Config = Configuration{}

type Configuration struct {
	Address      string
	ReadTimeout  int64
	WriteTimeOut int64
	Static       string
}

func P(a ...interface{}) {
	fmt.Println(a)
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
