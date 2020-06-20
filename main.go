package main

import (
	"fmt"
	"intercom/customer"
	"intercom/helper"
	"sort"
)

func main() {
	// Get lines frorm the URL
	var InputCustomers []string
	InputCustomers = helper.GetLinesFromURL("https://s3.amazonaws.com/intercom-take-home-test/customers.txt")

	var InvitedCustomers []helper.InvitedCustomer
	for _, line := range InputCustomers {
		// Unmarshal JSON for each line and
		var j helper.JSONCustomer
		j, err := helper.GetJSONCustomer(line)
		if err != nil {
			fmt.Printf("Bad Json [%s]\n", line)
			continue
		}
		// check whether customer is whithin 100km radius from the office
		var c customer.Customer
		c = *customer.NewCustomer(j.Name, j.UserID, j.Latt, j.Long)
		if c.IsInvitedForBeer() {
			i := helper.InvitedCustomer{c.UserID, c.Name}
			InvitedCustomers = append(InvitedCustomers, i)
		}
	}

	// Sort invited customers by userID and Print to stdout
	sort.Slice(InvitedCustomers, func(i, j int) bool {
		return InvitedCustomers[i].UserID < InvitedCustomers[j].UserID
	})
	for _, c := range InvitedCustomers {
		fmt.Printf("%#v\n", c)
	}
}
