package credentials

import "fmt"

type Credentials map[string]string

// GetCredentials returns map of user credentials for initiating API calls
func GetCredentials() Credentials {
	// The following variables aren't being pushed to Git.
	// Set them up as global constants in separate file when working on
	// a different computer.
	m := Credentials{
		"pAPIUserKey":    apiUserKey,
		"pEmail":         email,
		"pPassword":      password,
		"eGmailAddress":  gmailAddress,
		"eGmailPassword": gmailPassword,
	}
	return m
}

func (c Credentials) PrintInfo() {
	for k, v := range c {
		fmt.Printf("%v: %v\n", k, v)
	}
}
