package main

import "fmt"

// HOw do you know which input or file we need is in which package?
// Google it or simple read in the GOlang Documentaion

func main(){
	var conferenceName = "Rohit's Crowd Work"
	const conferenceTickets = 50
	var remainingTickets = 50

	// fmt.Println("Welcome to ", conferenceName, "booking Application")
	// Printf instead of 'ln'
	fmt.Printf("Welcome to %v booking application \n",conferenceName)
	fmt.Println("We have a total of", conferenceTickets,"tickets, of which only ",remainingTickets, " are remaining")
	fmt.Println("Hurry up! & Get your tickets here")


}