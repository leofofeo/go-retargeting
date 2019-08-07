package helpers

import "fmt"

// Columns is a []string type
type columns []string

// GetColumnsToKeep returns list of columns to keep in CSV file
func GetColumnsToKeep() columns {
	var columns = []string{"Prospect Id", "First Name", "Last Name", "Email", "Company", "Last Pageview Date", "Site Id", "First Assigned",
		"Score", "Sales Channel", "Annual Revenue", "Frequency", "Monetary", "Recency", "Industry",
		"Seniority", "Employee Range", "Estimated Annual Revenue", "Industry Group", "Sector", "Segment",
		"Sub Industry", "Job Title", "Department", "Company Type"}
	return columns
}

func (c columns) PrintInfo() {
	for _, column := range c {
		fmt.Println(column)
	}
}
