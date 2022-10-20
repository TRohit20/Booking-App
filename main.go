package main

import "fmt"

// HOw do you know which input or file we need is in which package?
// Google it or simple read in the GOlang Documentaion

func main(){
	var conferenceName = "Rohit's Crowd Work"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	// An alternative i.e Syntactic Sugar, IT ONLY APPLIES TO VAR NOT CONST/CONSTANTS
	// conferenceName := "Rohit"

	// fmt.Println("Welcome to ", conferenceName, "booking Application")
	// Printf instead of 'ln'
	fmt.Printf("Welcome to %v booking application \n",conferenceName)
	fmt.Println("We have a total of", conferenceTickets,"tickets, of which only ",remainingTickets, " are remaining")
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
	fmt.Println("Enter your First name:")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your E-mail:")
	fmt.Scan(&email)

	fmt.Println("Enter how many tickets you'd like to purchase:")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets \n",firstName,lastName,userTickets)
	fmt.Printf("You will receive a confirmation mail to your mail ID %v \n",email)

	remainingTickets -= userTickets
	fmt.Printf("only %v tickets remaining \n",remainingTickets)

}