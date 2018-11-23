package models

import "strings"

// statusMessages used to store all status Messages based on status Code
var statusMessages = map[string]string{
	"success":        "Success",
	"deletedsuccess": "Item deleted successfully",
	"notfound":       "IceCreams  not found",
	"insertsuccess":  "Ice Cream created successfully",
	"updatesuccess":  "Ice Cream updated successfully",
}

// GetStatusMessages used to get the status messages
func GetStatusMessages(value string) string {
	return statusMessages[strings.ToLower(value)]
}
