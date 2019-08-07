package credentials

import "fmt"

// Credentials map holds auth info necessary for authentication and subsequenet API calls
type Credentials map[string]string

// GetCredentials returns map of user credentials for initiating API calls
func GetCredentials() Credentials {
	// The following variables aren't being pushed to Git.
	// Set them up as global constants in separate file when working on
	// a different computer. Recommend creating said file in the same directory as credentials.go
	m := Credentials{
		"pAPIUserKey":    apiUserKey,
		"pEmail":         email,
		"pPassword":      password,
		"eGmailAddress":  gmailAddress,
		"eGmailPassword": gmailPassword,
	}
	return m
}

// PrintInfo of credentials to verify you're receiving the right info
func (c Credentials) PrintInfo() {
	for k, v := range c {
		fmt.Printf("%v: %v\n", k, v)
	}
}
