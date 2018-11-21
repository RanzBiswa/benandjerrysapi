package orders

import (
	"testing"

	"github.com/zalora_icecream/commonFramework/external/github.com/robfig/config"
)

func TestGetCustomerOrderStatusType(t *testing.T) {
	//set up
	c, e := config.ReadDefault("../crateAPI.cfg")
	if e != nil {
		t.Errorf("Error reading config file")
	}

	expectedType := "Open"

	typ := GetCustomerOrderStatusType(100, c)

	if typ != expectedType {
		t.Errorf("Wrong order type returned;should be %s, but is %s", expectedType, typ)
	}

	expectedType = "Cancelled"

	typ = GetCustomerOrderStatusType(900, c)

	if typ != expectedType {
		t.Errorf("Wrong order type returned;should be %s, but is %s", expectedType, typ)
	}
}
