package main

import (
	"strings"
)

func userInValid(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	validName := len(firstName) >= 2 && len(lastName) >= 2
	validMail := strings.Contains(email, "@")
	validTickets := userTickets > 0 && userTickets <= remainingTickets
	return validName, validMail, validTickets
}
