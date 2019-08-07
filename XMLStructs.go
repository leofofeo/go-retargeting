package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var apiKeyXML = []byte(`
<rsp stat="ok" version="1.0">
   <api_key>e82ca0f3f004e8291e6c634804296a06</api_key>
   <version>4</version>
</rsp>
`)

type rsp struct {
	Stat    string `xml:"type,attr"`
	APIKey  string `xml:api_key`
	version string `xml:version`
}

func getAPIKey(c credentials) string {
	baseURL := "https://pi.pardot.com/api/login/version/3 HTTP/1.1?"
	userKey := c["pAPIUserKey"]
	email := c["pEmail"]
	password := c["pPassword"]
	url := baseURL + "user_key=" + userKey + "&email=" + email + "&password=" + password
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))
	xmlResponseBody := string(body)
	xmlResponseBodyBytes := []byte(xmlResponseBody)
	fmt.Println("xmlResponseBodyBytes:", xmlResponseBodyBytes)
	xmlRSP := rsp{}
	xml.Unmarshal(xmlResponseBodyBytes, &xmlRSP)
	fmt.Println(xmlRSP)
	apiKey := "test"
	return apiKey
}
