package main

import "fmt"

// HOw do you know which input or file we need is in which package?
// Google it or simple read in the GOlang Documentaion

func main(){
	var conferenceName = "Rohit's Crowd Work"
	const conferenceTickets = 50
	var remainingTickets = 50
	// An alternative i.e Syntactic Sugar, IT ONLY APPLIES TO VAR NOT CONST/CONSTANTS
	// conferenceName := "Rohit"

	// fmt.Println("Welcome to ", conferenceName, "booking Application")
	// Printf instead of 'ln'
	fmt.Printf("Welcome to %v booking application \n",conferenceName)
	fmt.Println("We have a total of", conferenceTickets,"tickets, of which only ",remainingTickets, " are remaining")
	fmt.Println("Hurry up! & Get your tickets here")

	// Decalartion
	var userName string
	var userTickets int
	// Ask fro user input or Initialisation
	// userName = "Tom"
	// userTickets = 2

	// Taking a input, We do that using scan func and pointer 
	fmt.Scan(&userName)
	fmt.Scan(&userTickets)
	fmt.Printf("The user %v booked %v tickets. \n",	userName,userTickets)

}