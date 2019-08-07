package models

// ProspectRow captures all information about a prospect needed to export as single row in csv file
// ProspectRow should receive information from appropriate XMLStructs, which themselves unmarshal data from XMLResponse
type ProspectRow struct {
	ProspectID       string
	FirstName        string
	LastName         string
	Email            string
	Company          string
	LastPageViewDate string
	SiteID           string
	FirstAssigned    string
	Score            string
	SalesChannel     string
	AnnualRevenue    string
	Frequency        string
	Monetary         string
	Recency          string
	Industry         string
	EmployeeRange    string
	Segment          string
	JobTitle         string
	Department       string
	TotalPageViews   string
	MinutesOnSite    string
}
