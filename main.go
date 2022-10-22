package main

import (
	"fmt"
	"strings"
)

// HOw do you know which input or file we need is in which package?
// Google it or simple read in the GOlang Documentaion

func main() {
	var conferenceName = "Rohit's Crowd Work"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	// An alternative i.e Syntactic Sugar, IT ONLY APPLIES TO VAR NOT CONST/CONSTANTS
	// conferenceName := "Rohit"

	// var bookings [50]string //Array

	var bookings []string //Slices

	// Calling a function
	greet(conferenceName, conferenceTickets, int(remainingTickets))

	// fmt.Println("Welcome to ", conferenceName, "booking Application")
	// Printf instead of 'ln'
	// fmt.Printf("Welcome to %v booking application \n", conferenceName)
	// Taking a input, We do that using scan func and pointer

	for remainingTickets > 0 || len(bookings) < 50 {

		firstName, lastName, email, userTickets := takeInput()
		validName, validMail, validTickets := userInValid(firstName, lastName, email, int(userTickets), int(remainingTickets))

		if validMail && validName && validTickets {

			bookTicket(bookings, firstName, lastName, remainingTickets, userTickets, email)

			// Calling a function to go through the bookings array to print names of bookings
			firstNames := bNames(bookings)
			fmt.Printf("These are the First names of the bookings: %v \n", firstNames)

			// Boolean expression
			// var soldOut bool = remainingTickets == 0
			soldOut := remainingTickets == 0
			if soldOut {
				fmt.Println("Sorry, The tickets for the conference are all sold. Better luck next time")
				break
			}
		} else {
			if !validName {
				fmt.Println("Please enter a valid name ")
			}
			if !validMail {
				fmt.Println("Please provide a valid email ID")
			}
			if !validTickets {
				fmt.Println("PLease enter positive number of tickets")
			}
			continue
		}
	}

	// bookings[0] = firstName + " " + lastName
	// fmt.Printf("The elements in the Slice: %v \n",bookings)
	// fmt.Printf("The first value in the slice is: %v \n",bookings[0])
	// fmt.Printf("The length of slice is: %v \n",len(bookings))
	// fmt.Printf("Type of the Slice being: %T \n",bookings)

	// ** Switch cases
	// var city string
	// fmt.Println("Enter your city")
	// fmt.Scan(&city)

	// switch city {
	// case "Hyderabad":
	// 	fmt.Println("You choose hyderbad")
	// case "Bengalore", "Mumbai":
	// 	fmt.Println("Costly metropoliean cities")
	// case "Delhi":
	// 	fmt.Println("Capital of India")
	// default:
	// 	fmt.Println("Please choose a valid city")
	// }

}

func greet(confName string, conftickets int, remTicks int) {
	fmt.Printf("Welcome to %v's Ticket booking application \n", confName)
	fmt.Println("We have a total of", conftickets, "tickets, of which only ", remTicks, " are remaining")
	fmt.Println("Hurry up! & Get your tickets here")
}

func takeInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("Enter your First name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your E-mail:")
	fmt.Scan(&email)
	fmt.Println("Enter how many tickets you'd like to purchase:")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func bNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		// var firstName = names[0]
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func userInValid(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {
	validName := len(firstName) >= 2 && len(lastName) >= 2
	validMail := strings.Contains(email, "@")
	validTickets := userTickets > 0 && userTickets <= remainingTickets
	return validName, validMail, validTickets
}

func bookTicket(bookings []string, firstName string, lastName string, remainingTickets uint, userTickets uint, email string) {
	bookings = append(bookings, firstName+" "+lastName)
	remainingTickets -= userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets \n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation mail to your mail ID %v \n", email)
	fmt.Printf("only %v tickets remaining \n", remainingTickets)
}
