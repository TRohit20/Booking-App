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

	// fmt.Println("Welcome to ", conferenceName, "booking Application")
	// Printf instead of 'ln'
	fmt.Printf("Welcome to %v booking application \n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets, of which only ", remainingTickets, " are remaining")
	fmt.Println("Hurry up! & Get your tickets here")

	// Decalartion
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	// Ask fro user input or Initialisation
	// userName = "Tom"
	// userTickets = 2

	// Taking a input, We do that using scan func and pointer

	for remainingTickets > 0 || len(bookings) < 50 {
		fmt.Println("Enter your First name:")
		fmt.Scan(&firstName)
		fmt.Println("Enter your Last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your E-mail:")
		fmt.Scan(&email)

		fmt.Println("Enter how many tickets you'd like to purchase:")
		fmt.Scan(&userTickets)

		validName := len(firstName) >= 2 && len(lastName) >= 2
		validMail := strings.Contains(email, "@")
		validTickets := userTickets > 0 && userTickets <= remainingTickets

		if validMail && validName && validTickets {

			fmt.Printf("Thank you %v %v for booking %v tickets \n", firstName, lastName, userTickets)
			fmt.Printf("You will receive a confirmation mail to your mail ID %v \n", email)

			bookings = append(bookings, firstName+" "+lastName)

			remainingTickets -= userTickets

			fmt.Printf("only %v tickets remaining \n", remainingTickets)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				// var firstName = names[0]
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("These are the First names of the bookings: %v \n", firstNames)

			// Boolean expression
			// var soldOut bool = remainingTickets == 0
			soldOut := remainingTickets == 0
			if soldOut {
				fmt.Println("Sorry, The tickets for the conference are all sold. Better luck next time")
				break
			}
		} else {
			fmt.Println("You entered Invalid details, please check and re-enter")
			continue
		}
	}

	// bookings[0] = firstName + " " + lastName
	// fmt.Printf("The elements in the Slice: %v \n",bookings)
	// fmt.Printf("The first value in the slice is: %v \n",bookings[0])
	// fmt.Printf("The length of slice is: %v \n",len(bookings))
	// fmt.Printf("Type of the Slice being: %T \n",bookings)

}
