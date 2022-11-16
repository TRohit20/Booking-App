package main

import (
	"fmt"
	"sync"
	"time"
	// "strconv"
	// "strings"
)

// Package level variables declaration
var conferenceName = "Rohit's Crowd Work"

const conferenceTickets = 50

var RemainingTickets uint = 50

// var bookings = make(map[string]string, 0)
var bookings = make([]userData, 0)

type userData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets uint
}

// HOw do you know which input or file we need is in which package?
// Google it or simple read in the GOlang Documentaion

var syncing = sync.WaitGroup{}

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

	firstName, lastName, email, userTickets := takeInput()
	validName, validMail, validTickets := UserInValid(firstName, lastName, email, userTickets, RemainingTickets)

	if validMail && validName && validTickets {

		bookTicket(firstName, lastName, userTickets, email)

		syncing.Add(1)
		go sendTicket(userTickets, firstName, lastName, email)

		// Calling a function to go through the bookings array to print names of bookings
		firstNames := bNames()
		fmt.Printf("These are the First names of the bookings: %v \n", firstNames)

		// Boolean expression
		// var soldOut bool = remainingTickets == 0
		soldOut := RemainingTickets == 0
		if soldOut {
			fmt.Println("Sorry, The tickets for the conference are all sold. Better luck next time")

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
		syncing.Wait()
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
	fmt.Println("We have a total of", conferenceTickets, "tickets, of which only ", RemainingTickets, " are remaining")
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

		// var names = strings.Fields(booking)
		// var firstName = names[0]
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTicket(firstName string, lastName string, userTickets uint, email string) {
	// Creating a MAP for a user
	var userData = userData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		noOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("LIST OF ALL BOOKINGS IS %v \n", bookings)
	RemainingTickets -= userTickets

	fmt.Printf("Thank you %v %v for booking %v tickets \n", firstName, lastName, userTickets)
	fmt.Printf("You will receive a confirmation mail to your mail ID %v \n", email)
	fmt.Printf("only %v tickets remaining \n", RemainingTickets)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets have been booked under the name %v %v", userTickets, firstName, lastName)
	fmt.Println("******")
	fmt.Printf("Booking confirmed, Sending ticket: %v confirmation to %v", ticket, email)
	fmt.Println("******")
	syncing.Done()
}
