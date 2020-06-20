package helper

// JSONCustomer has name, userID and geo co-ordinates
type JSONCustomer struct {
	Name   string `json:"name"`
	UserID int    `json:"user_id"`
	Latt   string `json:"latitude"`
	Long   string `json:"longitude"`
}

// InvitedCustomer hold CustomerID and CustomerName
type InvitedCustomer struct {
	UserID       int
	CustomerName string
}
