package main

import (
	"fmt"
	"strings"
)

// Package level variables declaration
var conferenceName = "Rohit's Crowd Work"

const conferenceTickets = 50

var remainingTickets uint = 50
var bookings []string //Slices

// HOw do you know which input or file we need is in which package?
// Google it or simple read in the GOlang Documentaion

func main() {

	// An alternative i.e Syntactic Sugar, IT ONLY APPLIES TO VAR NOT CONST/CONSTANTS
	// conferenceName := "Rohit"

	// var bookings [50]string //Array

	// Calling a function
	greet()

	// fmt.Println("Welcome to ", conferenceName, "booking Application")
	// Printf instead of 'ln'
	// fmt.Printf("Welcome to %v booking application \n", conferenceName)
	// Taking a input, We do that using scan func and pointer

	for remainingTickets > 0 || len(bookings) < 50 {

		firstName, lastName, email, userTickets := takeInput()
		validName, validMail, validTickets := userInValid(firstName, lastName, email, userTickets)

		if validMail && validName && validTickets {

			bookTicket(firstName, lastName, userTickets, email)

			// Calling a function to go through the bookings array to print names of bookings
			firstNames := bNames()
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

func greet() {
	fmt.Printf("Welcome to %v's Ticket booking application \n", conferenceName)
	fmt.Println("We have a total of", conferenceTickets, "tickets, of which only ", remainingTickets, " are remaining")
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

func bNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		// var firstName = names[0]
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func bookTicket(firstName string, lastName string, userTickets uint, email string) {
	bookings = append(bookings, firstName+" "+lastName)
	remainingTickets -= userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets \n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation mail to your mail ID %v \n", email)
	fmt.Printf("only %v tickets remaining \n", remainingTickets)
}
