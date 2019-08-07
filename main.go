package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	credentials := getCredentials()
	apiKey := getAPIKey(credentials)
	makeVisitorAPICall(apiKey, credentials["pAPIUserKey"])
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

func makeVisitorAPICall(apiKey string, userKey string) {
	baseURL := "https://pi.pardot.com/api/prospect/version/4/do/query?"
	createdAfter := "last_7_days"
	createdBefore := "today"
	onlyIdentified := "true"

	securityParams := "user_key=" + userKey + "&api_key=" + apiKey
	filterParams := "&created_after" + createdAfter + "&created_before=" + createdBefore + "&only_identified=" + onlyIdentified
	url := baseURL + securityParams + filterParams

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))

}
