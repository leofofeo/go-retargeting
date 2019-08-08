package helpers

import (
	"fmt"
	"go-retargeting/models"
	"strings"
)

// Prospects is slice of Prospect Row
type Prospects []models.ProspectRow

// PrintSliceSize returns length of Prospects array, i.e., how many ProspectRow structs we're working with, and thus how many rows we should have
func (p Prospects) PrintSliceSize() {
	fmt.Println("Number of rows (prospects slice size):", len(p))
}

// OrganizeProspects creates a slice of Prospect structs with the data from the VisitorsXMLResp
func OrganizeProspects(xmlVisitors models.VisitorsXMLResp) Prospects {
	prospects := Prospects{}

	visitors := xmlVisitors.Result.Prospects
	for _, v := range visitors {

		if emailIsInvalid(v.Email) {
			// fmt.Printf("%v is an invalid email\n", v.Email)
			continue
		}
		pr := models.ProspectRow{}
		pr.ProspectID = v.ID
		pr.FirstName = v.FirstName
		pr.LastName = v.LastName
		pr.Email = v.Email
		pr.Company = v.Company
		pr.SiteID = v.SiteID
		pr.Score = v.Score
		pr.SalesChannel = v.SalesChannel
		pr.AnnualRevenue = v.AnnualRevenue
		pr.Frequency = v.Frequency
		pr.Monetary = v.Monetary
		pr.Recency = v.Recency
		pr.Industry = v.Industry
		pr.EmployeeRange = v.Employees
		pr.Segment = v.Segment
		pr.JobTitle = v.JobTitle
		pr.Department = v.Department

		prospects = append(prospects, pr)
	}

	return prospects
}

func emailIsInvalid(email string) bool {
	if strings.Contains(email, "epromos") || strings.Contains(email, "motivators") {
		return true
	}
	return false
}
