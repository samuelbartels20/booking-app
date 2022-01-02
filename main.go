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
	var userTickets int
	fmt.Println("Please enter your first name: ")
	fmt.Scan(&firstName)

	userTickets = 2
	fmt.Printf("user %v booked %v tickets.\n", userName, userTickets)

	
}