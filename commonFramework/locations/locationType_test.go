package locations

import (
	"testing"

	"github.com/zalora_icecream/commonFramework/external/github.com/robfig/config"
)

func TestTypeDesc(t *testing.T) {
	//set up
	c, e := config.ReadDefault("../crateAPI.cfg")
	if e != nil {
		t.Errorf("Error reading config file")
	}

	testGetLocationTypeDesc(t, c)
	testGetWarehouseTypeDesc(t, c)
	testGetBusinessTypeDesc(t, c)
}

func testGetLocationTypeDesc(t *testing.T, c *config.Config) {
	expectedDesc := "Home Store"

	desc := GetLocationTypeDesc("H", c)

	if desc != expectedDesc {
		t.Errorf("Wrong location type description returned;should be %s, but is %s", expectedDesc, desc)
	}

	expectedDesc = "Crossdock"

	desc = GetLocationTypeDesc("x", c)

	if desc != expectedDesc {
		t.Errorf("Wrong location type description returned;should be %s, but is %s", expectedDesc, desc)
	}

	expectedDesc = "Vendor Drop Ship"

	desc = GetLocationTypeDesc("y", c)

	if desc != expectedDesc {
		t.Errorf("Wrong location type description returned;should be %s, but is %s", expectedDesc, desc)
	}
}

func testGetWarehouseTypeDesc(t *testing.T, c *config.Config) {
	expectedDesc := "Warehouse"

	desc := GetWarehouseTypeDesc("100", c)

	if desc != expectedDesc {
		t.Errorf("Wrong warehouse type description returned;should be %s, but is %s", expectedDesc, desc)
	}

	expectedDesc = "Crossdock"

	desc = GetWarehouseTypeDesc("200", c)

	if desc != expectedDesc {
		t.Errorf("Wrong warehouse type description returned;should be %s, but is %s", expectedDesc, desc)
	}

	expectedDesc = "Store"

	desc = GetWarehouseTypeDesc("300", c)

	if desc != expectedDesc {
		t.Errorf("Wrong warehouse type description returned;should be %s, but is %s", expectedDesc, desc)
	}

	expectedDesc = "Hybrid"

	desc = GetWarehouseTypeDesc("400", c)

	if desc != expectedDesc {
		t.Errorf("Wrong warehouse type description returned;should be %s, but is %s", expectedDesc, desc)
	}
}

func testGetBusinessTypeDesc(t *testing.T, c *config.Config) {
	expectedDesc := "Central Furniture Warehouse"

	desc := GetBusinessTypeDesc("4", c)

	if desc != expectedDesc {
		t.Errorf("Wrong business type description returned;should be %s, but is %s", expectedDesc, desc)
	}

	expectedDesc = "Crossdock"

	desc = GetBusinessTypeDesc("998", c)

	if desc != expectedDesc {
		t.Errorf("Wrong business type description returned;should be %s, but is %s", expectedDesc, desc)
	}

	expectedDesc = "Western Housewares Warehouse"

	desc = GetBusinessTypeDesc("2", c)

	if desc != expectedDesc {
		t.Errorf("Wrong business type description returned;should be %s, but is %s", expectedDesc, desc)
	}

	expectedDesc = "Not a Warehouse"

	desc = GetBusinessTypeDesc("0", c)

	if desc != expectedDesc {
		t.Errorf("Wrong business type description returned;should be %s, but is %s", expectedDesc, desc)
	}
}
