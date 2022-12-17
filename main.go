package main

import "fmt"

func main() {
	const availableTickets = 50
	var remaningTickets = 50
	var conferenceName = "Go Conference"
	var username string
	var userTickets int

	fmt.Printf("Welcome to our Conference %v \n", conferenceName)
	fmt.Printf("The number of available tickets is %v  and we have %v remaining tickets\n", availableTickets, remaningTickets)
	fmt.Print("Book your ticket quickly !!\n")

	
	userTickets = 3
	username = "syrine"

	fmt.Printf("The user %v has %v tickets.\n", username, userTickets)

}