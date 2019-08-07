package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

type xmlResp struct {
	XMLName     xml.Name `xml:"rsp"`
	Description string   `xml:",innerxml"`
	APIKey      string   `xml:"api_key"`
	Version     string   `xml:"version"`
}

func getAPIKey(c credentials) string {
	baseURL := "https://pi.pardot.com/api/login/version/3 HTTP/1.1?"
	userKey := c["pAPIUserKey"]
	email := c["pEmail"]
	password := c["pPassword"]
	url := baseURL + "user_key=" + userKey + "&email=" + email + "&password=" + password

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	xmlResponseBodyBytes := body
	xmlRSP := xmlResp{}
	xml.Unmarshal(xmlResponseBodyBytes, &xmlRSP)
	apiKey := xmlRSP.APIKey
	return apiKey
}
