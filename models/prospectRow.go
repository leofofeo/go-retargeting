package models

// ProspectRow captures all information about a prospect needed to export as single row in csv file
// ProspectRow should receive information from appropriate XMLStructs, which themselves unmarshal data from XMLResponse
type ProspectRow struct {
	prospectID             string
	firstName              string
	lastName               string
	email                  string
	company                string
	lastPageViewDate       string
	siteID                 string
	firstAssigned          string
	score                  string
	salesChannel           string
	annualRevenue          string
	frequency              string
	monetary               string
	recency                string
	industry               string
	seniority              string
	employeeRange          string
	estimatedAnnualRevenue string
	industryGroup          string
	sector                 string
	segment                string
	subIndustry            string
	jobTitle               string
	department             string
	companyType            string
	totalPageViews         string
	minutesOnSite          string
}
