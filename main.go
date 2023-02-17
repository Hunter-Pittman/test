package main

import (
	"fmt"
	"os"
)

func main() {
	var eventName = "CISO Super Cyber Summit"
	const totalTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome to the", eventName, "event!")
	fmt.Printf("There are a total of %v tickets available.\n", totalTickets)
	fmt.Printf("There are a remainder of %v tickets.\n", remainingTickets)

	infoGathering(remainingTickets)
}

func infoGathering(remainingTickets int) {
	var (
		id          string
		firstName   string
		lastName    string
		email       string
		userTickets int
	)

	fmt.Println("Please enter your student ID:")
	fmt.Scan(&id)

	fmt.Println("Please enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email:")
	fmt.Scan(&email)

	fmt.Println("Please enter the number of tickets you wish to purchase:")
	fmt.Scan(&userTickets)

	if userTickets > remainingTickets {
		fmt.Println("There are not enough tickets available.")
		os.Exit(0)
	} else if userTickets == remainingTickets {
		fmt.Println("Congrats you got the last tickets!!!!!!!!!")
	} else {
		remainingTickets = remainingTickets - userTickets
		fmt.Printf("There are a remainder of %v tickets.\n", remainingTickets)
	}

	fmt.Println("##########Trasaction Complete##########")
	fmt.Println("Student ID:", id)
	fmt.Println("First Name:", firstName)
	fmt.Println("Last Name:", lastName)
	fmt.Println("Email:", email)
	fmt.Println("Number of Tickets:", userTickets)

}
