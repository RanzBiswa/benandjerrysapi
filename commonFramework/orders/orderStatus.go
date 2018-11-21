package orders

import (
	"strconv"

	"github.com/zalora_icecream/commonFramework/external/github.com/robfig/config"
)

//GetCustomerOrderStatusType Used to fetch customer order type based on the order status
func GetCustomerOrderStatusType(status int, c *config.Config) string {
	section := "customer_order_status"

	if !c.HasSection(section) {
		return ""
	}

	ordType, e := c.String(section, strconv.Itoa(status))

	if len(ordType) == 0 ||
		e != nil {
		return ""
	}

	return ordType

}
