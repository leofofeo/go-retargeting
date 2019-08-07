package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	// file := "07-29-2019_retargeting.csv"
	credentials := getCredentials()
	apiKey := getAPIKey(credentials)
	// makeVisitorAPICall()
	// readSegmentationFile(file)
	// columns := getColumnsToKeep()

	// columns.printInfo()
	// credentials.printInfo()
	fmt.Println("API key:", apiKey)
}

func readSegmentationFile(filename string) {
	csvFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("There seems to be a problem with your selected file. Please ensure it exists")
		os.Exit(1)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	fmt.Println(reader)
}

func makeVisitorAPICall() {
	// baseUrl := "https://pi.pardot.com/api/prospect/version/4/do/query?"
	// apiKey := ""
	// createdAfter := "last_7_days"
	// createdBefore := "today"
	// onlyIdentified := "true"
}
