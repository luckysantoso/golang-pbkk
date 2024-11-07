package main

import (
	"fmt"
	"strings"
)

func main() {

	const totalTickets int = 50
	var availableTickets uint = 50
	eventTitle := "Go Conference"
	reservations := []string{}

	fmt.Printf("Welcome to %v event registration system.\nWe have total of %v tickets and %v are still available.\nSecure your spot by booking tickets now\n", eventTitle, totalTickets, availableTickets)

	for {
		var fname string
		var lname string
		var email string
		var ticketCount uint

		// asking for user input
		fmt.Println("Provide Your First Name: ")
		fmt.Scanln(&fname)

		fmt.Println("Provide Your Last Name: ")
		fmt.Scanln(&lname)

		fmt.Println("Provide Your Email Address: ")
		fmt.Scanln(&email)

		fmt.Println("How many tickets do you need?: ")
		fmt.Scanln(&ticketCount)

		// validate user input
		nameIsValid := len(fname) >= 2 && len(lname) >= 2
		emailIsValid := strings.Contains(email, "@")
		ticketCountIsValid := ticketCount > 0 && ticketCount <= availableTickets

		if nameIsValid && emailIsValid && ticketCountIsValid {

			// book ticket in system
			availableTickets = availableTickets - ticketCount
			reservations = append(reservations, fname+" "+lname)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", fname, lname, ticketCount, email)
			fmt.Printf("%v tickets remaining for %v\n", availableTickets, eventTitle)

			// print only first names
			fnames := []string{}
			for _, booking := range reservations {
				var names = strings.Fields(booking)
				fnames = append(fnames, names[0])
			}
			fmt.Printf("List of attendees' first names: %v\n", fnames)

			// exit application if no tickets are left
			if availableTickets == 0 {
				// end program
				fmt.Println("All tickets sold out! See you at the next event.")
				break
			}
		} else {
			if !nameIsValid {
				fmt.Println("firt name or last name you entered is too short")
			}
			if !emailIsValid {
				fmt.Println("email address you entered doesn't contain @ sign")
			}
			if !ticketCountIsValid {
				fmt.Println("number of tickets you entered is invalid")
			}
			continue
		}
	}

	// Switch statement example
	city := "London"

	switch city {
	case "New York":
		// booking for New York conference
	case "Singapore", "Hong Kong":
		// booking for Singapore & Hong Kong conference
	case "London", "Berlin":
		// booking for London & Berlin conference
	case "Mexico City":
		// booking for Mexico City conference
	default:
		fmt.Print("No valid city selected")
	}
}
