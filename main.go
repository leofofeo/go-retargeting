package main

import (
	"go-retargeting/apicalls"
	"go-retargeting/credentials"
	"go-retargeting/helpers"
	"go-retargeting/models"
)

func main() {
	credentials := credentials.GetCredentials()
	apiKey := apicalls.GetAPIKey(credentials)
	prospects := helpers.Prospects{}
	visitorData := models.VisitorsXMLResp{}
	visitorData = apicalls.GetVisitorData(apiKey, credentials["pAPIUserKey"])
	visitorData.PrintSizeInfo()
	prospects = helpers.OrganizeProspects(visitorData)
	prospects.PrintSliceSize()
}
