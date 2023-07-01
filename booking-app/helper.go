package main

import "strings"

func validateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidemail := strings.Contains(email, "@")
	isValidTicketNumbers := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidemail, isValidTicketNumbers
}