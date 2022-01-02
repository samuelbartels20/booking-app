package main

import "fmt"

func main() {

	ConferenceName := "Go Conference"
	const conferenceTickets = 50
	remainingTickets := 50

	fmt.Printf("Welcome to %v booking application", ConferenceName)
	fmt.Printf("We have total of", conferenceTickets , "tickets and", remainingTickets, "are still available")
	fmt.Printf("Get your tickets here to attend")

	
}