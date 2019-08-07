package main

import (
	"go-retargeting/apicalls"
	"go-retargeting/credentials"
	"go-retargeting/helpers"
)

func main() {
	credentials := credentials.GetCredentials()
	apiKey := apicalls.GetAPIKey(credentials)
	visitorData := apicalls.GetVisitorData(apiKey, credentials["pAPIUserKey"])
	visitorData.PrintSizeInfo()
	prospects := helpers.OrganizeProspects(visitorData)
	prospects.PrintSliceSize()
}
