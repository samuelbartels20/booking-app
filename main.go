package main

import "fmt"

func main() {

	conferenceName := "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v, are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email: ")
	fmt.Scan(&email)

	fmt.Println("How many tickets are you buying? ")
	fmt.Scan(&userTickets)


	fmt.Printf("user %v booked %v tickets.\n", firstName, userTickets)

	
}