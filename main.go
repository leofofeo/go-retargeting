package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file := "07-29-2019_retargeting.csv"
	readSegmentationFile(file)
	columns := getColumnsToKeep()
	credentials := getCredentials()
	columns.printInfo()
	credentials.printInfo()
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
