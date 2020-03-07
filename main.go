package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const interval = 5
const configFilePath = "config.txt"
const logFilePath = "website-healthcheck.log"

func main() {

	// log setup
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()  // defers the execution until surroundings complete
	log.SetOutput(logFile) // comment this line to see logs in the console
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile | log.Lmsgprefix)

	printMenu()
	userInput := getUserInput()

	// menu handling
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
	//implicit var declaration
	appVersion := 0.1
	fmt.Println("Website monitoring tool")
	fmt.Println("Version:", appVersion)
	fmt.Println("1 - start monitor")
	fmt.Println("2 - view logs")
	fmt.Println("0 - exit")
}

// set user input into a variable and return
func getUserInput() int {
	var userInput int
	fmt.Scan(&userInput)
	fmt.Println("Selected command:", userInput)
	return userInput
}

// start the monitoring loop
func startMonitoring() {
	log.Println("Monitoring...")
	urlList := readConfigFile()
	for {
		for _, url := range urlList {
			testURL(url)
		}
		time.Sleep(interval * time.Second)
	}
}

// access the url and log the status code
func testURL(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(url, err)
	} else {
		log.Println(url, resp.StatusCode)
	}
}

// read file with the list of sites
func readConfigFile() []string {
	log.Println("Reading config.txt file...")

	// open file
	configFile, err := os.Open(configFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	//read file
	var urlList []string
	reader := bufio.NewReader(configFile)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		urlList = append(urlList, line)
	}
	configFile.Close()
	return urlList
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
