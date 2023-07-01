package main

import (
	"fmt"
	"strings"
)

const ConferenceTickets int = 50

var ConferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidemail, isValidTicketNumbers := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidemail && isValidTicketNumbers {

			bookTicket(userTickets, firstName, lastName, email)

			// call function print first name
			firstName := getFirstName()
			fmt.Printf("The first names of bookings are: %v\n", firstName)

			if remainingTickets == 0 {
				//end the program

				fmt.Println("Our conference is booked out")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("frst name or last name you entered is too short")
			}
			if !isValidemail {
				fmt.Println("email address you entered is too short")
			}
			if !isValidTicketNumbers {
				fmt.Println("number of tickets you entered is too short")
			}
			fmt.Println("Your Input data is invalid Try again ")
			// fmt.Printf("We have only %v tickets remaining, so you can't book %v ticlets\n", remainingTickets, userTickets)
		}

	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", ConferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available\n", ConferenceTickets, remainingTickets)
	fmt.Println("Get your tickerts here")
}

func getFirstName() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}

	return firstNames
}



func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//ask user about their data
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email name: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets, You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, ConferenceName)
}
