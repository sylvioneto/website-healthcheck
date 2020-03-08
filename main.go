package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

const configFilePath = "config.yaml"
const logFilePath = "website-healthcheck.log"

var config Config

// Config is the map used to unmarshall the config.yaml file
type Config struct {
	Interval   int      `yaml:"interval"`
	Websites   []string `yaml:",flow"`
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
		log.SetOutput(logFile)
	}

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)
	log.SetPrefix("website-healthcheck: ")

	// menu handling
	printMenu()
	userInput := getUserInput()
	switch userInput {
	case 1:
		startMonitoring()
	case 2:
		printLog()
	case 0:
		log.Println("Exit...")
		os.Exit(0)
	default:
		log.Fatalln("Not valid option")
		os.Exit(1)
	}
}

// print menu
func printMenu() {
	appVersion := 0.1
	fmt.Println("Website monitoring tool")
	fmt.Println("Version:", appVersion)
	fmt.Println("1 - start monitor")
	fmt.Println("2 - view logs")
	fmt.Println("0 - exit")
}

// request user input
func getUserInput() int {
	var userInput int
	fmt.Scan(&userInput)
	fmt.Println("Selected command:", userInput)
	return userInput
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

// print log file
func printLog() {
	log.Println("Viewing logs...")
	logFile, err := ioutil.ReadFile(logFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(logFile))
}
