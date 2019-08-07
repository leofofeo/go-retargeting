package models

import (
	"encoding/xml"
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
	Result  []Result `xml:"result"`
	Version string   `xml:"version,attr"`
}

// Result tag for Pardot response - contained within VisitorsXMLResp
type Result struct {
	XMLName      xml.Name   `xml:"result"`
	TotalResults string     `xml:"total_results"`
	Prospect     []Prospect `xml:"prospect"`
}

// Prospect tag for Pardot response, contained within the first result tag
type Prospect struct {
	XMLName      xml.Name     `xml:"prospect"`
	ID           string       `xml:"id"`
	CampaignID   string       `xml:"campaign_id"`
	Salutation   string       `xml:"salutation"`
	FirstName    string       `xml:"first_name"`
	LastName     string       `xml:"last_name"`
	Email        string       `xml:"email"`
	Address      string       `xml:"address_one"`
	City         string       `xml:"city"`
	State        string       `xml:"MO"`
	ZipCode      string       `xml:"zip"`
	PhoneNumber  string       `xml:"phone"`
	Source       string       `xml:"source"`
	Industry     string       `xml:"industry"`
	Score        string       `xml:"score"`
	DoNotEmail   bool         `xml:"is_do_not_email"`
	LastActivity lastActivity `xml:"last_activity"`
	Frequency    string       `xml:"Frequency"`
	LeadStatus   string       `xml:"Lead_Status"`
	Monetary     string       `xml:"Monetary"`
	Nurtured     string       `xml:"Nurtured"`
	SalesChannel string       `xml:"Sales_Channel"`
	Segment      string       `xml:"Segment"`
	SiteID       string       `xml:"Site_ID"`
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
