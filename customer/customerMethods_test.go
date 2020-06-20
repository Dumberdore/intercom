package customer

import (
	"fmt"
	"testing"
)

// TestIsInvitedForBeer case for known customers
func TestIsInvitedForBeer(t *testing.T) {

	var tests = []struct {
		customer Customer
		want     bool
	}{
		{*NewCustomer("Tang Cafe", 1, "53.340261", "-6.258568"), true},        // Tang Cafe - really nearby
		{*NewCustomer("Penneys WexFord", 2, "52.337682", "-6.460931"), false}, // Pennys, Wexford - more than 100km
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("Customer %10s is Invited [%t]", tt.customer.Name, tt.want)
		t.Run(testname, func(t *testing.T) {
			ans := tt.customer.IsInvitedForBeer()
			if ans != tt.want {
				t.Errorf("got %t, want %t", ans, tt.want)
			}
		})
	}
}
