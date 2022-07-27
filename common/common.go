package common

import "strings"

// Functions need to be exported to be made available in main.go
// To do this, simply capitalize the first letter of the function
// Variables are the same way. Capitalize the first letter to export it / make it global
func ValidateUserInput(firstName string, lastName string, email string, userTickets uint8, remainingTickets uint8) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".com")
	isValidTickets := userTickets <= remainingTickets && userTickets > 0
	return isValidName, isValidEmail, isValidTickets
}
