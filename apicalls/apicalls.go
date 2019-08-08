package apicalls

import (
	"encoding/xml"
	"fmt"
	"go-retargeting/credentials"
	"go-retargeting/models"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// GetAPIKey authenticates user info with api_user_key and gets api_key from Pardot API
func GetAPIKey(c credentials.Credentials) string {
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
	xmlRSP := models.AuthenticatingXMLResp{}
	xml.Unmarshal(xmlResponseBodyBytes, &xmlRSP)
	fmt.Println(xmlRSP)
	fmt.Println("xmlRSP api key is", xmlRSP.APIKey)
	fmt.Println("xmlRSP version is", xmlRSP.Version)
	fmt.Println("xmlRSP version tag is", xmlRSP.VersionTag)

	apiKey := xmlRSP.APIKey
	return apiKey
}

func makeVisitorAPICall(apiKey string, userKey string, offset int) []byte {
	baseURL := "https://pi.pardot.com/api/prospect/version/4/do/query?"
	createdAfter := "last_7_days"
	createdBefore := "today"
	onlyIdentified := "true"
	stringOffset := strconv.Itoa(offset)
	securityParams := "user_key=" + userKey + "&api_key=" + apiKey
	filterParams := "&created_after" + createdAfter + "&created_before=" + createdBefore + "&only_identified=" + onlyIdentified + "&offset=" + stringOffset
	url := baseURL + securityParams + filterParams

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

// GetVisitorData retrieves visitors in 7 day timeframe from Pardot API
func GetVisitorData(apiKey string, userKey string, offset int) models.VisitorsXMLResp {
	visitorData := makeVisitorAPICall(apiKey, userKey, offset)
	parsedVisitorData := parseVisitorData(visitorData)
	return parsedVisitorData
}

func parseVisitorData(b []byte) models.VisitorsXMLResp {
	visitorsXMLResp := models.VisitorsXMLResp{}
	fmt.Println(visitorsXMLResp)
	xml.Unmarshal(b, &visitorsXMLResp)
	return visitorsXMLResp
}
