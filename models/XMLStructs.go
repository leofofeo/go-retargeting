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
