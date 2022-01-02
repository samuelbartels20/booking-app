package main

import "fmt"

func main() {

	ConferenceName := "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50

	fmt.Println("Welcome to", ConferenceName , "booking application")
	fmt.Println("We have total of", conferenceTickets , "tickets and", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")

	
}