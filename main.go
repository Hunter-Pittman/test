package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

var remainingTickets uint8 = 50

type Add struct {
	human     string
	companion string
}

type Booking struct {
	id          string
	firstName   string
	lastName    string
	email       string
	userTickets uint8
	additional  []Add
}

func main() {
	var eventName = "CISO Super Cyber Summit"
	const totalTickets = 50

	fmt.Println("Welcome to the", eventName, "event!")
	fmt.Printf("There are a total of %v tickets available.\n", totalTickets)

	//fixed size array in go. This will show how many elements the array will hold.
	//var bookings = [50]string{"Hunter", "Ivan", "Ethan"}
	// ----------------------
	//You can also start adding and accessing elements based by their index (position)
	//var bookings [50]string
	//bookings[0] = "Hunter"
	//bookings[1] = "Ivan"
	//----------------------
	//Adding booking variable to the rest of the variables in the program
	//athough its usually better to use slices as it is more efficient and
	//and if you dont know how many objects are going to be in the array, a slice will
	//append new objects as they are added.
	//so thats why a slice would be better for our use case at least.
	for {
		infoGathering()
	}
}

func infoGathering() {
	var (
		id          string
		firstName   string
		lastName    string
		email       string
		userTickets uint8
		bookings    []Booking
	)

	// var (
	// 	id          = 5464
	// 	firstName   = "Hunter"
	// 	lastName    = "Harris"
	// 	email       = "hab@gmail.com"
	// 	userTickets = 1
	// 	bookings    []string
	// )

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

	//bookings = append(bookings, firstName+" "+lastName)

	booking := Booking{id: id, firstName: firstName, lastName: lastName, email: email, userTickets: userTickets}

	bookings = append(bookings, booking)

	fmt.Println("The whole slice: %v\n", bookings)

	// fmt.Println("##########Trasaction Complete##########")
	// fmt.Println("Student ID:", id)
	// fmt.Println("First Name:", firstName)
	// fmt.Println("Last Name:", lastName)
	// fmt.Println("Email:", email)
	// fmt.Println("Number of Tickets:", userTickets)
	// fmt.Println("##########Trasaction Complete##########")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"ID", "First Name", "Last Name", "Email", "Tickets"})

	for _, v := range bookings {
		var newUserTickets = string(v.userTickets)

		table.Append([]string{string(v.id), v.firstName, v.lastName, v.email, newUserTickets})
	}

	table.Render()
}

// fmt.Println("The whole slice: %v\n", bookings)
// fmt.Println("The first value: %v\n", bookings[0])
// //fmt.Println("Slice type: %T\n", bookings)
// fmt.Println("Slice length: %v\n", len(bookings))
