package customer

import "strconv"

// Customer has name, userID and geo co-ordinates
type Customer struct {
	Name   string
	UserID int
	Latt   float64
	Long   float64
}

// NewCustomer is wrapper to generate new customer
func NewCustomer(name string, userID int, latt string, long string) *Customer {
	lat1, _ := strconv.ParseFloat(latt, 64)
	lon1, _ := strconv.ParseFloat(long, 64)
	return &Customer{
		name,
		userID,
		lat1,
		lon1,
	}
}
