package locations

import (
	"strings"

	"github.com/zalora_icecream/commonFramework/external/github.com/robfig/config"
)

//GetLocationTypeDesc Used to fetch location type description based on the location type
func GetLocationTypeDesc(locType string, c *config.Config) string {
	return fetchConfigValue("location_type", locType, c)
}

//GetWarehouseTypeDesc Used to fetch warehouse type description based on the warehouse type
func GetWarehouseTypeDesc(warehouseType string, c *config.Config) string {
	return fetchConfigValue("warehouse_type", warehouseType, c)
}

//GetBusinessTypeDesc Used to fetch business type description based on the business type
func GetBusinessTypeDesc(busType string, c *config.Config) string {
	return fetchConfigValue("business_type", busType, c)
}

func fetchConfigValue(section string, key string, c *config.Config) string {

	if !c.HasSection(section) {
		return ""
	}

	value, e := c.String(section, strings.ToUpper(key))

	if len(value) == 0 ||
		e != nil {
		return ""
	}

	return value
}
