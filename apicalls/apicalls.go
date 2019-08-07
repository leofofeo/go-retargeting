package apicalls

import (
	"encoding/xml"
	"fmt"
	"go-retargeting/credentials"
	"go-retargeting/models"
	"io/ioutil"
	"log"
	"net/http"
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