package main

import "fmt"

type credentials map[string]string

func getCredentials() credentials {
	// The following variables aren't being pushed to Git.
	// Set them up as global constants in separate file when working on
	// a different computer.
	m := credentials{
		"pAPIUserKey":    apiUserKey,
		"pEmail":         email,
		"pPassword":      password,
		"eGmailAddress":  gmailAddress,
		"eGmailPassword": gmailPassword,
	}
	return m
}

func (c credentials) printInfo() {
	for k, v := range c {
		fmt.Printf("%v: %v\n", k, v)
	}
}
