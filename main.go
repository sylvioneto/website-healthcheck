package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

const configFilePath = "config.yaml"
const logFilePath = "website-healthcheck.log"
const logPrefix = "website-healthcheck: "

var config Config

// Config is the map used to unmarshall the config.yaml file
type Config struct {
	Interval   int      `yaml:"interval"`
	Websites   []string `yaml:"websites"`
	LogConsole bool     `yaml:"logConsole"`
}

func main() {

	readConfigFile()

	// set log file in case logConsole is set as false
	if !config.LogConsole {
		logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			log.Fatal(err)
		}
		defer logFile.Close()
		log.Println("Running... Log output ->", logFilePath)
		log.SetOutput(logFile)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)
	log.SetPrefix(logPrefix)

	startMonitoring()
}

// start health check monitor
func startMonitoring() {
	log.Println("Monitoring...")
	for {
		for _, url := range config.Websites {
			checkURL(url)
		}
		time.Sleep(time.Duration(config.Interval) * time.Second)
	}
}

// access the url and log the status code
func checkURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(url, err)
	} else {
		log.Println(url, resp.StatusCode)
	}
}

// read file with the list of sites
func readConfigFile() {
	log.Println("Reading config file...")
	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	// unmarshall yaml
	config = Config{}
	yaml.Unmarshal(configFile, &config)
	log.Println("Config:", config)
}
