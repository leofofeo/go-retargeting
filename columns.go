package main

import "fmt"

type columns []string

func getColumnsToKeep() columns {
	var columns = []string{"Prospect Id", "First Name", "Last Name", "Email", "Company", "Last Pageview Date", "Site Id", "First Assigned",
		"Score", "Sales Channel", "Annual Revenue", "Frequency", "Monetary", "Recency", "Industry",
		"Seniority", "Employee Range", "Estimated Annual Revenue", "Industry Group", "Sector", "Segment",
		"Sub Industry", "Job Title", "Department", "Company Type"}
	return columns
}

func (c columns) printInfo() {
	for _, column := range c {
		fmt.Println(column)
	}
}
