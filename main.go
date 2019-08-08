package main

import (
	"fmt"
	"go-retargeting/apicalls"
	"go-retargeting/credentials"
	"go-retargeting/helpers"
	"go-retargeting/models"
	"strconv"
)

func main() {
	credentials := credentials.GetCredentials()
	apiKey := apicalls.GetAPIKey(credentials)
	prospects := helpers.Prospects{}
	offset := 0
	visitorData := models.VisitorsXMLResp{}
	visitorData = apicalls.GetVisitorData(apiKey, credentials["pAPIUserKey"], offset)
	totalSize, _ := strconv.Atoi(visitorData.Result.TotalResults)
	prospects = helpers.OrganizeProspects(visitorData)
	fmt.Println("About to enter for loop")
	for {
		fmt.Println("In for loop")
		offset = 200 + offset
		fmt.Println("totalSize is", totalSize)
		fmt.Println("Offset (o) is", offset)
		if offset > totalSize {
			break
		}
		visitorData = apicalls.GetVisitorData(apiKey, credentials["pAPIUserKey"], offset)
		p := helpers.OrganizeProspects(visitorData)
		prospects = append(prospects, p...)
		fmt.Println("End of for loop")
	}
	prospects.PrintSliceSize()
}
