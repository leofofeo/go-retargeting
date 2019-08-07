package models

import (
	"encoding/xml"
	"fmt"
)

// AuthenticatingXMLResp corresponds to the initial Pardot response when authenticating user info to receive API key
type AuthenticatingXMLResp struct {
	XMLName    xml.Name `xml:"rsp"`
	Version    string   `xml:"version,attr"`
	APIKey     string   `xml:"api_key"`
	VersionTag string   `xml:"version"`
}

// VisitorsXMLResp corresponds to the Pardot response for visitors from GetVisitorsData
type VisitorsXMLResp struct {
	XMLName xml.Name `xml:"rsp"`
	Result  Result   `xml:"result"`
	Version string   `xml:"version,attr"`
}

// PrintSizeInfo total results as contained in the XML tag as well as the size of the Prospect array
func (v VisitorsXMLResp) PrintSizeInfo() {
	fmt.Println("Prospects:", len(v.Result.Prospects))
	fmt.Print("Total results:", v.Result.TotalResults)
	v.printProspectSampleInfo(0)
}

func (v VisitorsXMLResp) printProspectSampleInfo(size int) {
	for i := size; i < size+20; i++ {
		p := v.Result.Prospects[i]
		p.printInfo()
	}
}

// Result tag for Pardot response - contained within VisitorsXMLResp
type Result struct {
	XMLName      xml.Name   `xml:"result"`
	TotalResults int        `xml:"total_results"`
	Prospects    []Prospect `xml:"prospect"`
}

// Prospect tag for Pardot response, contained within the first result tag
type Prospect struct {
	XMLName       xml.Name     `xml:"prospect"`
	ID            string       `xml:"id"`
	CampaignID    string       `xml:"campaign_id"`
	Salutation    string       `xml:"salutation"`
	FirstName     string       `xml:"first_name"`
	LastName      string       `xml:"last_name"`
	Email         string       `xml:"email"`
	Company       string       `xml:"company"`
	Address       string       `xml:"address_one"`
	City          string       `xml:"city"`
	State         string       `xml:"MO"`
	ZipCode       string       `xml:"zip"`
	PhoneNumber   string       `xml:"phone"`
	Source        string       `xml:"source"`
	Industry      string       `xml:"industry"`
	Score         string       `xml:"score"`
	DoNotEmail    bool         `xml:"is_do_not_email"`
	LastActivity  lastActivity `xml:"last_activity"`
	Frequency     string       `xml:"Frequency"`
	LeadStatus    string       `xml:"Lead_Status"`
	Monetary      string       `xml:"Monetary"`
	Nurtured      string       `xml:"Nurtured"`
	SalesChannel  string       `xml:"Sales_Channel"`
	Segment       string       `xml:"Segment"`
	SiteID        string       `xml:"Site_ID"`
	AnnualRevenue string       `xml:"annual_revenue"`
	Recency       string       `xml:"Recency"`
	Employees     string       `xml:"employees"`
	JobTitle      string       `xml:"job_title"`
	Department    string       `xml:"department"`
}

func (p Prospect) printInfo() {
	fmt.Println("------- New Prospect -------")
	fmt.Println("ID:", p.ID)
	fmt.Println("First Name:", p.FirstName)
	fmt.Println("Last Name:", p.LastName)
	fmt.Println("Email:", p.Email)
	fmt.Println("Source:", p.Source)
	fmt.Println("Frequency:", p.Frequency)
	fmt.Println("Score:", p.Score)
	fmt.Println("Lead Status:", p.LeadStatus)
	fmt.Println("Sales Channel:", p.SalesChannel)
	fmt.Println("Campaign Id:", p.CampaignID)
	fmt.Println("City:", p.City)
	fmt.Println("Address:", p.Address)
	fmt.Println("Do Not Email:", p.DoNotEmail)
	fmt.Println("Frequency:", p.Frequency)
	fmt.Println("Industry:", p.Industry)
	fmt.Println("Last Activity:", p.LastActivity)
	fmt.Println("Nurtured:", p.Nurtured)
	fmt.Println("Monetary:", p.Monetary)
	fmt.Println("Phone Number:", p.PhoneNumber)
	fmt.Println("State:", p.State)
	fmt.Println("Segment:", p.Segment)
	fmt.Println("Site ID:", p.SiteID)
	fmt.Println("-----------")
}

type lastActivity struct {
	XMLName         xml.Name          `xml:"last_activity"`
	VisitorActivity []visitorActivity `xml:"visitor_activity"`
}

type visitorActivity struct {
	XMLName    xml.Name `xml:"visitor_activity"`
	ID         string   `xml:"id"`
	VisitorID  string   `xml:"visitor_id"`
	ProspectID string   `xml:"prospect_id"`
	Type       string   `xml:"type"`
	TypeName   string   `xml:"type_name"`
	Details    string   `xml:"details"`
}
