package models

import "strings"

// statusMessages used to store all status Messages based on status Code
var statusMessages = map[string]string{
	"SUCCESS":  "Success",
	"DESTROY":  " deleted successfully",
	"NOTFOUND": " not found",
}

// GetStatusMessages used to get the status messages
func GetStatusMessages(value string) string {
	return statusMessages[strings.ToUpper(value)]
}
