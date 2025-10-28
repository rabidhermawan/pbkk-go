package main

import (
	"fmt"
	"strings"
)

func main() {
	// var conferenceName string = "Go Conference"
	conferenceName := "Go Conference" // This can't work if we are declaring constant or explicitly defining a type
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// var bookings = [50]string{"Nana", "Nicole", "Peter"}
	// bookings := []string{"Nana", "Nicole", "Peter"}

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// Ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}

		// for index, booking := range bookings {
		for _, booking := range bookings {
			names := strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("There are all our bookings: %v\n", firstNames)
	}

}
