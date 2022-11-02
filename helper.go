package main

import (
	"strings"
)

func UserInValid(firstName string, lastName string, email string, userTickets uint, RemainingTickets uint) (bool, bool, bool) {
	validName := len(firstName) >= 2 && len(lastName) >= 2
	validMail := strings.Contains(email, "@")
	validTickets := userTickets > 0 && userTickets <= RemainingTickets
	return validName, validMail, validTickets
}
