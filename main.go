package main

import (
	"go-retargeting/apicalls"
	"go-retargeting/credentials"
	"go-retargeting/helpers"
)

func main() {
	credentials := credentials.GetCredentials()
	columns := helpers.GetColumnsToKeep()
	columns.PrintInfo()
	apiKey := apicalls.GetAPIKey(credentials)
	visitorData := apicalls.GetVisitorData(apiKey, credentials["pAPIUserKey"])
	visitorData.PrintByteSliceSize()
}
